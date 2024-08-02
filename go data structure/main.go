package main

//arbol completo cuando todos tienen el maximo o 0 nodos hijo

import (
	"fmt"
	"time"

	"github.com/Unbot2313/controllers"
	"github.com/Unbot2313/models"
)

//arbol binario prueba

var RootNode models.TreeNode

func main() {

	StartTime := time.Now()

	controllers.InitArbol(&RootNode)

	controllers.GetOneElement(&RootNode, 8)

	controllers.GetOneElement(&RootNode, 9)

	controllers.DeleteNode(&RootNode, 9)

	controllers.GetOneElement(&RootNode, 9)

	TiempoTotal := time.Since(StartTime)

	fmt.Println("Tiempo total: " + TiempoTotal.String())
}
