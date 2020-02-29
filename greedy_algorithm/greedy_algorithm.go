package greedy_algorithm

func Compute(n, m int, distanceMatrix [][]float32) [] int {

	// There are no sets in go, I use bool as value because its light
	selectedSet := make(map[int] bool)

	firstSelected := firstSelected(n, distanceMatrix)

	selectedSet[firstSelected] = true

	notSelectedSet := generateNotSelectedSet(firstSelected, n, distanceMatrix)

	for len(selectedSet) < m {
		selected := selectDiverseFromSet(notSelectedSet)

		selectedSet[selected] = true
		delete(notSelectedSet, selected)

		recalculateMinDistances(notSelectedSet, distanceMatrix, selected)
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

func generateNotSelectedSet (firstSelected, n int, distanceMatrix [][]float32 ) map[int]float32 {

	notSelectedSet := make (map[int]float32)

	for i := 0 ; i < n ; i++ {
		if i != firstSelected {
			notSelectedSet[i] = distanceMatrix[i][firstSelected]
		}
	}

	return notSelectedSet
}

func selectDiverseFromSet (notSelectedSet map[int]float32)  int {

	var maxValue float32 = 0
	max := -1

	for notSelected, minDistance := range notSelectedSet {
		if max == -1 || minDistance > maxValue {
			maxValue = minDistance
			max = notSelected
		}
	}

	return max
}

func recalculateMinDistances (notSelectedSet map[int]float32, distanceMatrix [][]float32, selected int) {

	for notSelected, minDistance := range notSelectedSet {
		distanceWithNewSelected := distanceMatrix[notSelected][selected]
		if distanceWithNewSelected < minDistance {
			notSelectedSet[notSelected] = distanceWithNewSelected
		}
	}
}

func setToSlice (selectedSet map[int]bool) (slice []int) {
	slice = make([]int, 0, len(selectedSet))

	for elem := range selectedSet {
		slice = append(slice, elem)
	}

	return slice
}
