package greedy_algorithm

import (
	"math"
)

func Compute(n, m int, distanceMatrix [][]float32) [] int {

	// There are no sets in go, I use bool as value because its light
	selectedSet := make(map[int] bool)

	firstSelected := firstSelected(n, distanceMatrix)

	selectedSet[firstSelected] = true

	for len(selectedSet) < m {
		selected := selectDiverseFromSet(selectedSet, distanceMatrix, n)

		selectedSet[selected] = true
	}

	return setToSlice(selectedSet)

}

func firstSelected (n int, distanceMatrix [][]float32) (selected int) {

	var maxValue float32 = 0

	for i := 0 ; i < n ; i++ {
		var current float32 = 0
		for j := 0 ; j < n ; j++ {
			current += distanceMatrix[i][j]
		}
		if current > maxValue {
			maxValue = current
			selected = i
		}
	}

	return selected
}

func selectDiverseFromSet (selectedSet map[int]bool, distanceMatrix [][] float32, n int)  int {

	var maxValue float32 = 0
	max := 0

	for i := 0 ; i < n ; i++ {
		_, contains := selectedSet[i]
		if !contains {
			minDistance := float32(math.MaxFloat32)
			for alreadySelected := range selectedSet {
				if distanceMatrix[i][alreadySelected] < minDistance {
					minDistance = distanceMatrix[i][alreadySelected]
				}
			}
			if minDistance > maxValue {
				maxValue = minDistance
				max = i
			}
		}
	}

	return max
}

func setToSlice (selectedSet map[int]bool) (slice []int) {
	slice = make([]int, 0, len(selectedSet))

	for elem := range selectedSet {
		slice = append(slice, elem)
	}

	return slice
}
