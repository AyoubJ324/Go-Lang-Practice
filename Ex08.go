package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)
	exit := make(chan bool)

	go run(t1, chan1)
	go run(t2, chan2)

	var output bool

	go func() {
		for i := range chan1 {
			if i == <-chan2 {
				output = true
			} else {
				output = false
			}
		}
		exit <- true
	}()
	<-exit
	return output
}

func run(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func main() {
	fmt.Println("Generating trees . . .")
	fmt.Println()

	t1 := tree.New(10)
	t2 := tree.New(5)
	fmt.Println("-------TEST ONE------")
	fmt.Println("Tree 1: ", t1)
	fmt.Println("Tree 2: ", t2)
	fmt.Println("Comparing Tree 1 & Tree 2 . . .")
	fmt.Println("Test Result = ", Same(t1, t2))
	fmt.Println("---------END----------")

	fmt.Println()

	t3 := tree.New(5)
	t4 := tree.New(5)
	fmt.Println("-------TEST TWO------")
	fmt.Println("Tree 3: ", t3)
	fmt.Println("Tree 4: ", t4)
	fmt.Println("Comparing Tree 3 & Tree 4 . . .")
	fmt.Println("Test Result = ", Same(t3, t4))
	fmt.Println("---------END----------")

}
