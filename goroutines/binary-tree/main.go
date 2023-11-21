package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t Tree) String() string {
	return fmt.Sprintf("Left: %v\nValue: %v\nRight: %v", t.Left, t.Value, t.Right)
}

func main() {
	tree1 := Tree{Value: 1}
	tree2 := Tree{Value: 2}
	tree1.Left = &tree2
	tree3 := Tree{Value: 3}
	tree1.Right = &tree3
	tree4 := Tree{Value: 4}
	tree2.Left = &tree4
	tree5 := Tree{Value: 5}
	tree2.Right = &tree5
	tree6 := Tree{Value: 6}
	tree4.Left = &tree6

	fmt.Println(tree1)
	fmt.Println(tree2)
	fmt.Println(tree3)
	fmt.Println(tree4)
	fmt.Println(tree5)
	fmt.Println(tree6)
}
