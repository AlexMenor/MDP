package ils_algorithm

import (
	"MDP/local_search_algorithm"
	"math/rand"
)

func Compute(n, m int, distanceMatrix [][]float32) []int {
	t := m / 10

	sol, _ := local_search_algorithm.ComputeForBMB(n, m, distanceMatrix)

	bestDiversity := getDiversity(sol, distanceMatrix, m)
	bestSol := make([]int, m)
	copy(bestSol, sol)


	const MAX_ITERATIONS int = 9

	for i := 0 ; i < MAX_ITERATIONS ; i++ {
		mutate(sol, n, t)
		sol, _ = local_search_algorithm.ComputeForILS(n,m, distanceMatrix, sol)
		currentDiversity := getDiversity(sol, distanceMatrix, m)

		if currentDiversity > bestDiversity {
			bestDiversity = currentDiversity
			copy(bestSol, sol)
		} else {
			copy(sol, bestSol)
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

func mutate(selected[]int, n int, t int){
	notSelected := getNotSelected(selected,n)

	rand.Shuffle(len(selected), func(i, j int) { selected[i], selected[j] = selected[j], selected[i]})
	rand.Shuffle(len(notSelected), func(i, j int) { notSelected[i], notSelected[j] = notSelected[j], notSelected[i]})

	for i := 0 ; i < t ; i++{
		selected[i] = notSelected[i]
	}

}

func getNotSelected(selected[]int, n int)[]int {
	notSelected := make([]int, 0,n - len(selected))

	selectedSet := make(map[int]bool)

	for _, s := range selected {
		selectedSet[s] = true
	}

	for i := 0 ; i < n ; i++ {
		_, isSelected := selectedSet[i]
		if !isSelected {
			notSelected = append(notSelected,i)
		}
	}

	return notSelected
}
