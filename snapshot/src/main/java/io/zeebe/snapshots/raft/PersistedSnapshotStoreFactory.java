/*
 * Copyright Camunda Services GmbH and/or licensed to Camunda Services GmbH under
 * one or more contributor license agreements. See the NOTICE file distributed
 * with this work for additional information regarding copyright ownership.
 * Licensed under the Zeebe Community License 1.0. You may not use this file
 * except in compliance with the Zeebe Community License 1.0.
 */
package io.zeebe.snapshots.raft;

import java.nio.file.Path;

/**
 * Creates a snapshot store which should store its {@link PersistedSnapshot} and {@link
 * TransientSnapshot} instances in the given directory.
 */
@FunctionalInterface
public interface PersistedSnapshotStoreFactory {

  /**
   * Creates a snapshot store operating in the given {@code directory}.
   *
   * @param directory the root directory where snapshots should be stored
   * @param partitionName the partition name for this store
   * @return a new {@link PersistedSnapshotStore}
   */
  PersistedSnapshotStore createSnapshotStore(Path directory, String partitionName);
}
