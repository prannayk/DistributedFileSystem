package main

import (
	"github.com/prannayk/DistributedFileSystem/src/node"
	"sync"
)
	
type Node struct {
	ChordStructure.Node
	Mutex sync.RWMutex
	FingerTable fingerTable
}

func (node *Node) FindSuccessor(nextHash []byte) (*Node, error) {
	// TODO(prannayk) : unimplemented method
	return node, nil;
}

func NewNode(id []byte, Name string) *Node {
	n := &Node {}
	n.Id = id
	n.Name = Name
	return n
}
