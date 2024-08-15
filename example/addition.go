package main

import (
	"flag"
	"fmt"
	mlgo "ml-go"
)

func main() {
	var epochs int
	flag.IntVar(&epochs, "epochs", 10, "number of epochs [default: 10]")
	flag.Parse()

	n := mlgo.Network()
	n.AddLayer(mlgo.Linear(2, 1))

	for i := 0; i < epochs; i++ {
		fmt.Printf("epoch %d\n", i+1)

		x := mlgo.Matrix([][]float64{{1}, {1}})
		pred := n.Forward(x)
		y := mlgo.Matrix([][]float64{{2}})
		fmt.Printf("input: %v; pred: %v; actual: %v\n", x, pred, y)

		loss := mlgo.MSE(y, pred)
		fmt.Printf("loss: %v\n", loss)

		n.Backward(loss)
		n.Update(mlgo.LEARNING_RATE)
		fmt.Println()
	}
	fmt.Println("==============================")

	x := mlgo.Matrix([][]float64{{1}, {1}})
	y := n.Forward(x)
	fmt.Printf("input: %v; output: %v;\n", x, y)
}
