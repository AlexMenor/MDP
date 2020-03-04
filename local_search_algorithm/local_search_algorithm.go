package local_search_algorithm

import (
	"math/rand"
)

func Compute(n, m int, distanceMatrix [][]float32, MAX_ITERATIONS int) []int {
	selected, notSelected := generateRandomSets(n, m)
	computeContribution(selected, distanceMatrix)
	iterations := 0

	listOfMinContributors := getListOfMinContributors(selected)

	for i := 0 ; i < len(listOfMinContributors) ; i++ {
		minContributor := listOfMinContributors[i].key
		minContribution := listOfMinContributors[i].contribution

		candidate, candidateContribution, improvesSolution := tryCandidates(selected, notSelected, minContributor, minContribution, distanceMatrix)

		if improvesSolution {
			updateSets(selected, notSelected, minContributor, candidate, candidateContribution)
			recalculateContributions(selected, candidate, minContributor, distanceMatrix)
			listOfMinContributors = getListOfMinContributors(selected)
			i = 0
		}

		iterations++
		if iterations == MAX_ITERATIONS {
			return setAsSlice(selected, m)
		}
	}

	return setAsSlice(selected, m)
}

func updateSets (selected map[int]float32, notSelected map[int]bool, minContributor int, candidate int, candidateContribution float32) {
	delete(selected, minContributor)
	selected[candidate] = candidateContribution
	delete(notSelected, candidate)
	notSelected[minContributor] = true
}

func recalculateContributions(selected map[int]float32, added int, removed int, distanceMatrix [][]float32) {
	for selectedElem := range selected {
		if selectedElem != added {
			selected[selectedElem] -= distanceMatrix[selectedElem][removed]
			selected[selectedElem] += distanceMatrix[selectedElem][added]
		}
	}
}

func tryCandidates(selected map[int]float32, notSelected map[int]bool, minContributor int, minContribution float32, distanceMatrix [][]float32) (int, float32, bool) {
	for candidate := range notSelected {
		candidateContribution := computeCandidateContribution(candidate, minContributor, selected, distanceMatrix)

		increment := candidateContribution - minContribution

		if increment > 0 {
			return candidate, candidateContribution, true
		}
	}

	return 0, 0, false
}

func computeCandidateContribution(candidate, excluded int, selected map[int]float32, distanceMatrix [][]float32) float32 {

	var candidateContribution float32 = 0

	for alreadySelected := range selected {
		if alreadySelected != excluded {
			candidateContribution += distanceMatrix[alreadySelected][candidate]
		}
	}

	return candidateContribution

}

func generateRandomSets(n int, m int) (map[int]float32, map[int]bool) {
	selected := make(map[int]float32)
	for len(selected) < m {
		newSelected := rand.Intn(n)
		selected[newSelected] = 0.0
	}

	notSelected := make(map[int]bool)

	for i := 0; i < n; i++ {
		_, alreadySelected := selected[i]

		if !alreadySelected {
			notSelected[i] = true
		}
	}

	return selected, notSelected
}

func computeContribution(selected map[int]float32, distanceMatrix [][]float32) {
	for i := range selected {
		var currentContribution float32 = 0
		for j := range selected {
			currentContribution += distanceMatrix[i][j]
		}
		selected[i] = currentContribution
	}
}

func setAsSlice(selected map[int]float32, m int) []int {
	sliceOfSelected := make([]int, 0, m)
	for selectedElem := range selected {
		sliceOfSelected = append(sliceOfSelected, selectedElem)
	}

	return sliceOfSelected
}
