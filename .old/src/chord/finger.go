package main

import (
	"fmt"
	"math/big"
	"bytes"
)

type fingerTable []*fingerEntry
//type Node  ChordStructure.Node
type fingerEntry struct {
	ID []byte 				// ID hash of the relevant entry
	RemoteNode *Node
}

func newFingerTable(node *Node) fingerTable {
	ft := make([]*fingerEntry, config.KeySize) 
	for i := range ft {
		ft[i] = newFingerEntry(fingerMath(node.Id, i), node)
	}
	// Completed initializing the fingerTable for the given node
	return ft
}

func newFingerEntry(startID []byte, RemoteNode* Node) *fingerEntry {
	return &fingerEntry{
		ID : startID,
		RemoteNode : RemoteNode,
	}
}

func fingerMath(NodeID []byte, index int) []byte {
	iInt := big.NewInt(2)
	iInt.Exp(iInt, big.NewInt(int64(index)), config.Max)
	mInt := big.NewInt(2)
	mInt.Exp(mInt, big.NewInt(int64(config.KeySize)), config.Max)
	res := &big.Int{}
	res.SetBytes(NodeID).Add(res, iInt).Mod(res, mInt)
	return PadID(res.Bytes())
}

func (ft fingerTable) String() string {				// converts fingerTable to string, testing purposes ONLY
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("FingerTable"))
	for _, val := range ft {
		buf.WriteString(fmt.Sprintf("\n\t{start:%v\tnodeID:%v %v}",
						IDToString(val.ID),
						IDToString(val.RemoteNode.Id),
						val.RemoteNode.Name,
						))
	}
	return buf.String()
}

func (node *Node) FingerTableString() string { 			// Go routine safe wrapper
	node.FTMutex.RLock()
	defer node.FTMutex.RUnlock()
	return node.FingerTable.String()
}

func (node *Node) FixNextFinger(next int) int {			// gets the nextfinger in ring, used for stabilization
	nextHash := fingerMath(node.Id, next)				// TODO(prannayk) : test this, untested presently
	succ, err := node.FindSuccessor(nextHash)			// implemented in chord / ring
	if err != nil {
		return next
	}
	finger := newFingerEntry(nextHash, succ)
	node.FTMutex.Lock()
	node.FingerTable[next] = finger
	node.FTMutex.Unlock()
	return (next + 1) % config.KeySize
}


// Testing functions 
func TestFinger () {
	InitConfigWrapper()
	t := []byte("789456789")
	t = PadID(t)
	fmt.Println(t)
	fmt.Println(fingerMath(t, 1))
	newNode := NewNode(t, "0.0.0.0")
	fTable := newFingerTable(newNode)
	newNode.FingerTable = fTable
	fmt.Println(fTable.String())
	fmt.Println(newNode.FingerTableString())
}

func main() {
	fmt.Println("Main is in finger.go, testing from here")
}
