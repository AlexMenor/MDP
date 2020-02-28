package main

import (
	"MDP/greedy_algorithm"
	"MDP/problem_reader"
	"fmt"
)

func main() {

	file := "problem_instances/MDG-b_1_n500_m50.txt"

	n, m, distanceMatrix := problem_reader.ReadFile(file)

	sol := greedy_algorithm.Compute(n,m, distanceMatrix)

	fmt.Println(sol)
	fmt.Println(getDiversity(sol, distanceMatrix, m))

}

func getDiversity(selected []int, distanceMatrix [][]float32, m int) (diversity float32) {

	for i := 0 ; i < m - 1 ; i++ {
		for j := i + 1 ; j < m ; j++ {
			diversity += distanceMatrix[selected[i]][selected[j]]
		}
	}

	return diversity
}
