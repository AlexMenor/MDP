package simulated_annealing_algorithm

import (
	"math"
	"math/rand"
)

func ComputeForILS(n, m int, distanceMatrix [][]float32, initialSolution []int) []int {

	maxEvaluations := 10000
	var alpha float32 = 0.8

	if len(initialSolution) != 0 {
		notSelected := getNotSelected(initialSolution, n)
		currentDiversity := getDiversity(initialSolution, distanceMatrix, m)

		return ComputeWithParameters(n, m, distanceMatrix, maxEvaluations, initialSolution, notSelected, currentDiversity, alpha)
	} else {
		selected, notSelected, currentDiversity := GenRandomSolution(n, m, distanceMatrix)

		return ComputeWithParameters(n, m, distanceMatrix, maxEvaluations, selected, notSelected, currentDiversity, alpha)

	}

}

func Compute(n, m int, distanceMatrix [][]float32) []int {
	selected, notSelected, currentDiversity := GenRandomSolution(n, m, distanceMatrix)

	return ComputeWithParameters(n, m, distanceMatrix, 100000, selected, notSelected, currentDiversity, 0.95)

}

func ComputeWithParameters(n, m int, distanceMatrix [][]float32, maxEvaluations int, selected []int, notSelected []int, currentDiversity float32, alpha float32) []int {
	maxNeighbours := n
	maxSuccess := maxNeighbours / 10
	annealings := maxEvaluations / maxNeighbours

	bestDiversity := currentDiversity
	bestSolution := make([]int, len(selected))
	copy(bestSolution, selected)

	//finalTemp := 0.001
	currentTemp := (currentDiversity * 0.3) / 1.204 // -Ln(0.3)
	//beta := (float32(currentTemp) - float32(finalTemp)) / (float32(annealings) * currentTemp * float32(finalTemp))

	for a := 0; a < annealings; a++ {
		currentNeighbours := 0
		currentSuccess := 0

		for currentNeighbours < maxNeighbours && currentSuccess < maxSuccess {
			i := rand.Intn(m)
			j := rand.Intn(n - m)
			randomSelected := selected[i]
			randomNotSelected := notSelected[j]
			diversityOfSelected := diversityOfOneElement(selected, randomSelected, randomSelected, distanceMatrix)
			diversityOfNotSelected := diversityOfOneElement(selected, randomSelected, randomNotSelected, distanceMatrix)

			increment := diversityOfNotSelected - diversityOfSelected

			probabilityOfChange := math.Exp(float64(increment / currentTemp))

			if increment > 0 || rand.Float64() <= probabilityOfChange {
				swapElements(i, j, selected, notSelected)
				currentDiversity += increment
				currentSuccess++

				if currentDiversity > bestDiversity {
					copy(bestSolution, selected)
					bestDiversity = currentDiversity
				}
			}
			currentNeighbours++

		}
		if currentSuccess == 0 {
			return bestSolution
		} else {
			//currentTemp = currentTemp / (1 + beta*currentTemp)
			currentTemp *= alpha
		}
	}
	return bestSolution
}

func GenRandomSolution(n, m int, distanceMatrix [][]float32) ([]int, []int, float32) {
	selected := make(map[int]bool)

	for len(selected) < m {
		randomSelected := rand.Intn(n)
		selected[randomSelected] = true
	}

	toReturnSelected := make([]int, 0, m)
	toReturnNotSelected := make([]int, 0, n-m)

	for i := 0; i < n; i++ {
		_, isSelected := selected[i]
		if isSelected {
			toReturnSelected = append(toReturnSelected, i)
		} else {
			toReturnNotSelected = append(toReturnNotSelected, i)
		}
	}

	return toReturnSelected, toReturnNotSelected, getDiversity(toReturnSelected, distanceMatrix, m)
}

func getDiversity(selected []int, distanceMatrix [][]float32, m int) (diversity float32) {

	for i := 0; i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			diversity += distanceMatrix[selected[i]][selected[j]]
		}
	}

	return diversity
}

func diversityOfOneElement(selected []int, excluded int, element int, distanceMatrix [][]float32) float32 {
	var toReturn float32 = 0

	for _, s := range selected {
		if s != excluded {
			toReturn += distanceMatrix[s][element]
		}
	}

	return toReturn
}

func swapElements(i, j int, selected, notSelected []int) {
	temp := selected[i]
	selected[i] = notSelected[j]
	notSelected[j] = temp
}

func getNotSelected(selected []int, n int) []int {
	notSelected := make([]int, 0, n-len(selected))

	selectedSet := make(map[int]bool)

	for _, s := range selected {
		selectedSet[s] = true
	}

	for i := 0; i < n; i++ {
		_, isSelected := selectedSet[i]
		if !isSelected {
			notSelected = append(notSelected, i)
		}
	}

	return notSelected
}
