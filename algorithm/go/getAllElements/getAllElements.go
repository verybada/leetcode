package getAllelements

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	valuesChan := make(chan int, 10)
	go func() {
		if root1 != nil {
			listTree(root1, valuesChan)
		}
		if root2 != nil {
			listTree(root2, valuesChan)
		}

		close(valuesChan)
	}()

	values := make([]int, 0)
	for {
		value, more := <-valuesChan
		if !more {
			break
		}

		values = append(values, value)

	}
	sort.Ints(values)
	return values
}

func listTree(root *TreeNode, valuesChan chan int) {
	if root.Left != nil {
		listTree(root.Left, valuesChan)
	}

	valuesChan <- root.Val

	if root.Right != nil {
		listTree(root.Right, valuesChan)
	}
}
