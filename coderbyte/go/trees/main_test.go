package main

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func generateTree() *Node {
	// create nodes
	n4 := NewLeaf('E')
	n3 := NewLeaf('D')
	n2 := NewLeaf('C')
	n1 := NewNode('B', n3, n4)
	root := NewNode('A', n1, n2)

	return root
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestPreOrder(t *testing.T) {
	tree := generateTree()
	in := []Data{}
	expected := []Data{'A', 'B', 'D', 'E', 'C'}

	out := tree.PreOrder(in)
	log.Debugf("PreOrder Result: %s", out)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("wanted %#v, got %#v", expected, out)
	}
}

func TestInOrder(t *testing.T) {
	tree := generateTree()
	in := []Data{}
	expected := []Data{'D', 'B', 'E', 'A', 'C'}

	out := tree.InOrder(in)
	log.Debugf("InOrder Result: %s", out)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("wanted %#v, got %#v", expected, out)
	}
}

func TestPostOrder(t *testing.T) {
	tree := generateTree()
	in := []Data{}
	expected := []Data{'D', 'E', 'B', 'C', 'A'}

	out := tree.PostOrder(in)
	log.Debugf("PostOrder Result: %s", out)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("wanted %s, got %s", expected, out)
	}
}

func TestLevelOrder(t *testing.T) {
	tree := generateTree()
	in := []Data{}
	expected := []Data{'A', 'B', 'C', 'D', 'E'}

	out := tree.LevelOrder(in)
	log.Debugf("LevelOrder Result: %s", out)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("wanted %s, got %s", expected, out)
	}
}
