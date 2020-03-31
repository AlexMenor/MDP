package common

import "math/rand"

func GenRandomPoblation(distanceMatrix [][]float32, n, m, tam int) []Chromosome {
	toReturn := make([]Chromosome, 0, tam)
	for i := 0; i < tam; i++ {
		newChromosome := GenRandomChromosome(distanceMatrix, m, n)
		toReturn = append(toReturn, newChromosome)
	}
	return toReturn
}

func GetSelectedFromPoblation(currentPoblation []Chromosome, selectedTam int) []Chromosome {
	toReturn := make([]Chromosome, 0, selectedTam)
	for i := 0; i < selectedTam; i++ {
		winner := getAWinner(currentPoblation)
		toReturn = append(toReturn, winner)
	}
	return toReturn
}

func PositionalCrossOver(father Chromosome, mother Chromosome, distanceMatrix [][]float32) (firstChild, secondChild Chromosome) {
	nGenes := len(father.genes)

	childGenes := make([]bool, nGenes, nGenes)
	restGenes := make(map[int]bool)

	for g := 0; g < nGenes; g++ {
		if father.genes[g] == mother.genes[g] {
			childGenes[g] = mother.genes[g]
		} else {
			restGenes[g] = mother.genes[g]
		}
	}

	firstChildGenes := getCompleteChildGenes(childGenes, restGenes)
	secondChildGenes := getCompleteChildGenes(childGenes, restGenes)

	firstChild = GenChromosomeFromGenes(firstChildGenes, distanceMatrix)
	secondChild = GenChromosomeFromGenes(secondChildGenes, distanceMatrix)

	return firstChild, secondChild
}

func getCompleteChildGenes(childGenes []bool, restGenes map[int]bool) []bool {
	nGenes := len(childGenes)
	toReturn := make([]bool, nGenes)

	shuffleGenes(restGenes)

	for i := 0; i < nGenes; i++ {
		value, unAsigned := restGenes[i]

		if unAsigned {
			toReturn[i] = value
		} else {
			toReturn[i] = childGenes[i]
		}
	}

	return toReturn
}

func shuffleGenes(genes map[int]bool) {
	rand.Shuffle(len(genes), func(i, j int) { genes[i], genes[j] = genes[j], genes[i] })
}

func getAWinner(currentPoblation []Chromosome) Chromosome {
	n := len(currentPoblation)
	firstIndex := rand.Intn(n)
	secondIndex := rand.Intn(n)

	for secondIndex != firstIndex {
		secondIndex = rand.Intn(n)
	}

	first := currentPoblation[firstIndex]
	second := currentPoblation[secondIndex]

	if first.value > second.value {
		return first
	} else {
		return second
	}
}
