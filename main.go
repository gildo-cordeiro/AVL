package main

import (
	"fmt"

	"github.com/gildo-cordeiro/AVL/avltree"
)

func main() {
	var root *avltree.Node

	// Exemplo de inserção de chaves na árvore AVL
	root = avltree.Insert(root, 10)
	root = avltree.Insert(root, 20)
	root = avltree.Insert(root, 30)
	root = avltree.Insert(root, 40)
	root = avltree.Insert(root, 50)
	root = avltree.Insert(root, 25)

	// Exibindo a árvore AVL em ordem
	fmt.Println("Árvore AVL em ordem:")
	avltree.InorderTraversal(root)
}
