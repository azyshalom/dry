package mocks

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

//SwarmDockerDaemon mocks a DockerDaemon operating in Swarm mode
type SwarmDockerDaemon struct {
	DockerDaemonMock
}

// Info provides a mock function with given fields:
func (_m *SwarmDockerDaemon) Info() (types.Info, error) {
	clusterInfo := swarm.ClusterInfo{ID: "MyClusterID"}
	swarmInfo := swarm.Info{
		LocalNodeState:   swarm.LocalNodeStateActive,
		NodeID:           "ThisNodeID",
		Cluster:          &clusterInfo,
		ControlAvailable: true}
	return types.Info{
		Name:     "test",
		NCPU:     2,
		MemTotal: 1024,
		Swarm:    swarmInfo}, nil
}

//Node returns a node with the given id
func (_m *SwarmDockerDaemon) Node(id string) (*swarm.Node, error) {
	return &swarm.Node{ID: id}, nil
}

//Nodes returns a list of nodes with 1 element
func (_m *SwarmDockerDaemon) Nodes() ([]swarm.Node, error) {
	return []swarm.Node{swarm.Node{ID: "1"}}, nil
}

//Tasks mock
func (_m *SwarmDockerDaemon) Tasks(nodeID string) ([]swarm.Task, error) {
	return []swarm.Task{swarm.Task{NodeID: nodeID}}, nil
}