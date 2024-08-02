package main

//arbol completo cuando todos tienen el maximo o 0 nodos hijo

import (
	"fmt"

	"github.com/Unbot2313/controllers"
	"github.com/Unbot2313/models"
)

//arbol binario prueba

var RootNode models.TreeNode

func main() {

	fmt.Println(RootNode)

	controllers.InitArbol(&RootNode)

	fmt.Println(RootNode)

	controllers.GetOneElement(&RootNode, 8)
	controllers.GetOneElement(&RootNode, 9)
}
