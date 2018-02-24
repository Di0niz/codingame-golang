package main

import (
	"fmt"
	"testing"
)

func TestCreateGraph(t *testing.T) {

	g := NewGraph()

	g.AppendLink(0, 1)
	g.AppendLink(0, 2)
	g.AppendLink(2, 3)
	g.AppendLink(1, 3)
	g.AppendLink(3, 5)
	g.AppendLink(3, 7)
	g.AppendLink(7, 8)
	g.AppendLink(8, 9)
	g.AppendLink(9, 6)
	g.AppendLink(5, 9)

	fmt.Println(g)

	for val := range g.Neighbors(1) {
		fmt.Println(val)
	}

	fmt.Println(g.Search(2, 6))

}

func TestSimpleTest(t *testing.T) {

	g := NewGraph()

	g.AppendLink(0, 1)
	g.AppendLink(0, 2)
	g.AppendLink(2, 3)
	g.AppendLink(1, 3)

	fmt.Println(g)

	fmt.Println(g.Search(2, 3))

}
