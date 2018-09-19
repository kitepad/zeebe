/*
 * Copyright © 2017 camunda services GmbH (info@camunda.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.zeebe.gateway;

import com.google.protobuf.Empty;
import io.zeebe.gateway.api.commands.PartitionInfo;
import io.zeebe.gateway.api.commands.Topology;
import io.zeebe.gateway.api.commands.Workflow;
import io.zeebe.gateway.api.events.DeploymentEvent;
import io.zeebe.gateway.api.events.JobEvent;
import io.zeebe.gateway.api.events.WorkflowInstanceEvent;
import io.zeebe.gateway.cmd.ClientException;
import io.zeebe.gateway.protocol.GatewayOuterClass.BrokerInfo;
import io.zeebe.gateway.protocol.GatewayOuterClass.BrokerInfo.Builder;
import io.zeebe.gateway.protocol.GatewayOuterClass.CreateJobResponse;
import io.zeebe.gateway.protocol.GatewayOuterClass.CreateWorkflowInstanceResponse;
import io.zeebe.gateway.protocol.GatewayOuterClass.DeployWorkflowResponse;
import io.zeebe.gateway.protocol.GatewayOuterClass.HealthResponse;
import io.zeebe.gateway.protocol.GatewayOuterClass.Partition;
import io.zeebe.gateway.protocol.GatewayOuterClass.Partition.PartitionBrokerRole;
import io.zeebe.gateway.protocol.GatewayOuterClass.WorkflowResponseObject;
import java.util.ArrayList;

public class ResponseMapper {

  private PartitionBrokerRole remapPartitionBrokerRoleEnum(
      final io.zeebe.gateway.api.commands.BrokerInfo brokerInfo, final PartitionInfo partition) {
    switch (partition.getRole()) {
      case LEADER:
        return PartitionBrokerRole.LEADER;
      case FOLLOWER:
        return PartitionBrokerRole.FOLLOW;
      default:
        throw new ClientException(
            "Unknown broker role in response for partition "
                + partition
                + " on broker "
                + brokerInfo);
    }
  }

  public HealthResponse toHealthResponse(final Topology brokerResponse) {
    final HealthResponse.Builder healthResponseBuilder = HealthResponse.newBuilder();
    final ArrayList<BrokerInfo> infos = new ArrayList<>();

    for (final io.zeebe.gateway.api.commands.BrokerInfo el : brokerResponse.getBrokers()) {
      final Builder brokerInfo = BrokerInfo.newBuilder();
      brokerInfo.setHost(el.getHost());
      brokerInfo.setPort(el.getPort());

      for (final PartitionInfo p : el.getPartitions()) {
        final Partition.Builder partitionBuilder = Partition.newBuilder();
        partitionBuilder.setPartitionId(p.getPartitionId());
        partitionBuilder.setRole(remapPartitionBrokerRoleEnum(el, p));
        brokerInfo.addPartitions(partitionBuilder);
      }

      infos.add(brokerInfo.build());
    }

    healthResponseBuilder.addAllBrokers(infos);
    return healthResponseBuilder.build();
  }

  public DeployWorkflowResponse toDeployWorkflowResponse(final DeploymentEvent brokerResponse) {
    final DeployWorkflowResponse.Builder deployWorkflowResponseBuilder =
        DeployWorkflowResponse.newBuilder();

    for (final Workflow workflow : brokerResponse.getWorkflows()) {
      deployWorkflowResponseBuilder.addWorkflows(
          WorkflowResponseObject.newBuilder()
              .setBpmnProcessId(workflow.getBpmnProcessId())
              .setVersion(workflow.getVersion())
              .setWorkflowKey(workflow.getWorkflowKey())
              .setResourceName(workflow.getResourceName()));
    }
    return deployWorkflowResponseBuilder.build();
  }

  public Empty emptyResponse(final Object response) {
    return Empty.getDefaultInstance();
  }

  public CreateJobResponse toCreateJobResponse(final JobEvent jobEvent) {
    return CreateJobResponse.newBuilder()
        .setKey(jobEvent.getKey())
        .setPartitionId(jobEvent.getMetadata().getPartitionId())
        .build();
  }

  public CreateWorkflowInstanceResponse toCreateWorkflowInstanceResponse(
      final WorkflowInstanceEvent workflowInstanceEvent) {
    return CreateWorkflowInstanceResponse.newBuilder()
        .setWorkflowKey(workflowInstanceEvent.getWorkflowKey())
        .setBpmnProcessId(workflowInstanceEvent.getBpmnProcessId())
        .setVersion(workflowInstanceEvent.getVersion())
        .setPartitionId(workflowInstanceEvent.getMetadata().getPartitionId())
        .setWorkflowInstanceKey(workflowInstanceEvent.getWorkflowInstanceKey())
        .build();
  }
}