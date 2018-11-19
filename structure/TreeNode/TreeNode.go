package TreeNode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTestData(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = string([]rune(data)[1 : len(data)-1])
	res := strings.Split(data, ",")
	length := len(res)
	treeNode := make([]TreeNode, length)
	for i := 0; i < length; i++ {
		if res[i] != "nil" {
			val, err := strconv.Atoi(res[i])
			if err != nil {
				panic(err)
			}
			treeNode[i] = TreeNode{val, nil, nil}
		}
	}
	for i := 0; i < length; i++ {
		if treeNode[i].Val != 0 {
			leftIndex := i*2 + 1
			if leftIndex < length && treeNode[leftIndex].Val != 0 {
				treeNode[i].Left = &treeNode[leftIndex]
			}
			rightIndex := leftIndex + 1
			if rightIndex < length && treeNode[leftIndex].Val != 0 {
				treeNode[i].Right = &treeNode[rightIndex]
			}
		}
	}
	return &treeNode[0]
}

const spce = "      "

func Print(root *TreeNode) {
	treeNodePrint(root, 0)
}

func treeNodePrint(node *TreeNode, deep int) {
	if node == nil {
		printSpace(deep)
		fmt.Println("#")
		return
	}
	treeNodePrint(node.Right, deep+1)
	printSpace(deep)
	printNode(node.Val)
	treeNodePrint(node.Left, deep+1)
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(spce)
	}
}

func printNode(val int) {
	var buffer strings.Builder
	temp := strconv.Itoa(val)
	buffer.WriteString(temp)
	buffer.WriteString("<")
	spceNum := len(spce) - buffer.Len()
	for i := 0; i < spceNum; i++ {
		buffer.WriteString(" ")
	}
	fmt.Println(buffer.String())
}
