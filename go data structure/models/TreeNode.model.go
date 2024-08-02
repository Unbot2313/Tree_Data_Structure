package models

type TreeNode struct {
	Valor int

	//Derecha son numeros mayores, Izquierda numeros menores
	Izquierda, Derecha *TreeNode
}
