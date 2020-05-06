package bmb_algorithm

import "MDP/local_search_algorithm"

func Compute(n, m int, distanceMatrix [][]float32) []int {
	const MAX_SEARCHES int = 10
	var bestDiversity float32 = 0
	bestSol := make([]int, m)

	for i := 0; i < MAX_SEARCHES; i++ {
		sol, _ := local_search_algorithm.ComputeForBMB(n, m, distanceMatrix)
		thisSolDiversity := getDiversity(sol, distanceMatrix, m)
		if thisSolDiversity > bestDiversity {
			bestDiversity = thisSolDiversity
			copy(bestSol, sol)
		}
	}

	return bestSol
}

func getDiversity(selected []int, distanceMatrix [][]float32, m int) (diversity float32) {

	for i := 0; i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			diversity += distanceMatrix[selected[i]][selected[j]]
		}
	}

	return diversity
}
