/*
 * Copyright Camunda Services GmbH and/or licensed to Camunda Services GmbH under
 * one or more contributor license agreements. See the NOTICE file distributed
 * with this work for additional information regarding copyright ownership.
 * Licensed under the Zeebe Community License 1.0. You may not use this file
 * except in compliance with the Zeebe Community License 1.0.
 */
package io.zeebe.broker.system.partitions.impl;

import static org.assertj.core.api.Assertions.assertThat;

import io.atomix.raft.snapshot.SnapshotChunk;
import io.atomix.raft.snapshot.TransientSnapshot;
import io.atomix.raft.zeebe.ZeebeEntry;
import io.atomix.storage.journal.Indexed;
import io.zeebe.broker.system.partitions.SnapshotReplication;
import io.zeebe.broker.system.partitions.snapshot.impl.FileBasedSnapshotStore;
import io.zeebe.broker.system.partitions.snapshot.impl.FileBasedSnapshotStoreFactory;
import io.zeebe.db.impl.DefaultColumnFamily;
import io.zeebe.db.impl.rocksdb.ZeebeRocksDbFactory;
import io.zeebe.test.util.AutoCloseableRule;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.function.Consumer;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.TemporaryFolder;

public final class FailingSnapshotChunkReplicationTest {

  @Rule public final TemporaryFolder tempFolderRule = new TemporaryFolder();
  @Rule public final AutoCloseableRule autoCloseableRule = new AutoCloseableRule();

  private StateControllerImpl replicatorSnapshotController;
  private StateControllerImpl receiverSnapshotController;
  private FileBasedSnapshotStore senderStore;
  private FileBasedSnapshotStore receiverStore;

  public void setup(final SnapshotReplication replicator) throws IOException {
    final var senderRoot = tempFolderRule.newFolder("sender").toPath();
    senderStore =
        (FileBasedSnapshotStore)
            new FileBasedSnapshotStoreFactory().createSnapshotStore(senderRoot, "1");

    final var receiverRoot = tempFolderRule.newFolder("receiver").toPath();
    receiverStore =
        (FileBasedSnapshotStore)
            new FileBasedSnapshotStoreFactory().createSnapshotStore(receiverRoot, "1");

    replicatorSnapshotController =
        new StateControllerImpl(
            1,
            ZeebeRocksDbFactory.newFactory(DefaultColumnFamily.class),
            senderStore,
            senderRoot.resolve("runtime"),
            replicator,
            l ->
                Optional.of(
                    new Indexed(l, new ZeebeEntry(1, System.currentTimeMillis(), 1, 10, null), 0)),
            db -> Long.MAX_VALUE);
    senderStore.addSnapshotListener(replicatorSnapshotController);

    receiverSnapshotController =
        new StateControllerImpl(
            1,
            ZeebeRocksDbFactory.newFactory(DefaultColumnFamily.class),
            receiverStore,
            receiverRoot.resolve("runtime"),
            replicator,
            l ->
                Optional.ofNullable(
                    new Indexed(l, new ZeebeEntry(1, System.currentTimeMillis(), 1, 10, null), 0)),
            db -> Long.MAX_VALUE);
    receiverStore.addSnapshotListener(receiverSnapshotController);

    autoCloseableRule.manage(replicatorSnapshotController);
    autoCloseableRule.manage(senderStore);
    autoCloseableRule.manage(receiverSnapshotController);
    autoCloseableRule.manage(receiverStore);
    replicatorSnapshotController.openDb();
  }

  @Test
  public void shouldNotWriteChunksAfterReceivingInvalidChunk() throws Exception {
    // given
    final EvilReplicator replicator = new EvilReplicator();
    setup(replicator);
    final var transientSnapshot = takeSnapshot();

    // when
    transientSnapshot.persist();

    // then
    final List<SnapshotChunk> replicatedChunks = replicator.replicatedChunks;
    assertThat(replicatedChunks.size()).isGreaterThan(0);

    assertThat(receiverStore.getLatestSnapshot()).isEmpty();
  }

  @Test
  public void shouldNotMarkSnapshotAsValidIfNotReceivedAllChunks() throws Exception {
    // given
    final FlakyReplicator replicator = new FlakyReplicator();
    setup(replicator);
    final var transientSnapshot = takeSnapshot();

    // when
    transientSnapshot.persist();

    // then
    final List<SnapshotChunk> replicatedChunks = replicator.replicatedChunks;
    assertThat(replicatedChunks.size()).isGreaterThan(0);
    assertThat(receiverStore.getLatestSnapshot()).isEmpty();
  }

  private TransientSnapshot takeSnapshot() {
    receiverSnapshotController.consumeReplicatedSnapshots();
    return replicatorSnapshotController.takeTransientSnapshot(1).orElseThrow();
  }

  private final class FlakyReplicator implements SnapshotReplication {

    final List<SnapshotChunk> replicatedChunks = new ArrayList<>();
    private Consumer<SnapshotChunk> chunkConsumer;

    @Override
    public void replicate(final SnapshotChunk snapshot) {
      replicatedChunks.add(snapshot);
      if (chunkConsumer != null) {
        if (replicatedChunks.size() < 3) {
          chunkConsumer.accept(snapshot);
        }
      }
    }

    @Override
    public void consume(final Consumer<SnapshotChunk> consumer) {
      chunkConsumer = consumer;
    }

    @Override
    public void close() {}
  }

  private final class EvilReplicator implements SnapshotReplication {

    final List<SnapshotChunk> replicatedChunks = new ArrayList<>();
    private Consumer<SnapshotChunk> chunkConsumer;

    @Override
    public void replicate(final SnapshotChunk snapshot) {
      replicatedChunks.add(snapshot);
      if (chunkConsumer != null) {
        chunkConsumer.accept(
            replicatedChunks.size() > 1 ? new DisruptedSnapshotChunk(snapshot) : snapshot);
      }
    }

    @Override
    public void consume(final Consumer<SnapshotChunk> consumer) {
      chunkConsumer = consumer;
    }

    @Override
    public void close() {}
  }

  private final class DisruptedSnapshotChunk implements SnapshotChunk {

    private final SnapshotChunk snapshotChunk;

    DisruptedSnapshotChunk(final SnapshotChunk snapshotChunk) {
      this.snapshotChunk = snapshotChunk;
    }

    @Override
    public String getSnapshotId() {
      return snapshotChunk.getSnapshotId();
    }

    @Override
    public int getTotalCount() {
      return snapshotChunk.getTotalCount();
    }

    @Override
    public String getChunkName() {
      return snapshotChunk.getChunkName();
    }

    @Override
    public long getChecksum() {
      return 0;
    }

    @Override
    public byte[] getContent() {
      return snapshotChunk.getContent();
    }

    @Override
    public long getSnapshotChecksum() {
      return snapshotChunk.getSnapshotChecksum();
    }
  }
}
