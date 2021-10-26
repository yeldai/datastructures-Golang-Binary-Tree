package main

import "fmt"


var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}



func (n *Node) Insert (k int){
	if n.Key <k{
		if n.Right ==nil{
			n.Right= &Node{Key:k}
		}else{
			n.Right.Insert(k)
		}
		
	}else if n.Key >k{
		if n.Left ==nil{
			n.Left =&Node{Key:k}
		}else{
			n.Left.Insert(k)
		}
		
	}
}


// Search


func (n *Node) Search (k int) bool{
	count++
	if n==nil{
		return false 
	}

	if n.Key <k{
	return n.Right.Search(k)
	}else if n.Key>k{
		return n.Left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)
	tree.Insert(203)
	tree.Insert(12)
	tree.Insert(76)
	tree.Insert(22)
	tree.Insert(32200)
	tree.Insert(2)
	tree.Insert(222)
	tree.Insert(3020)
	

	tree.Search(76)
	fmt.Println(tree.Search(76))
	fmt.Println(tree.Search(32200))
	fmt.Println(count)
}