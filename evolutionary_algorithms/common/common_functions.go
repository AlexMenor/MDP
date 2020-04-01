package common

import (
	"math/rand"
	"sort"
)

func GenRandomPoblation(distanceMatrix [][]float32, n, m, tam int) Poblation {
	toReturn := make(Poblation, 0, tam)
	for i := 0; i < tam; i++ {
		newChromosome := GenRandomChromosome(distanceMatrix, n, m)
		toReturn = append(toReturn, newChromosome)
	}
	return toReturn
}

func GetSelectedFromPoblation(currentPoblation Poblation) Poblation {
	selectedTam := len(currentPoblation)
	toReturn := make(Poblation, 0, selectedTam)
	for i := 0; i < selectedTam; i++ {
		winner := getAWinner(currentPoblation)
		toReturn = append(toReturn, winner)
	}
	return toReturn
}

func PositionalCrossover(father Chromosome, mother Chromosome, distanceMatrix [][]float32) (firstChild, secondChild Chromosome) {
	nGenes := len(father.genes)

	commonGenes := make([]bool, nGenes, nGenes)
	unasignedGenesIndex := make([]int, 0)
	unasignedGenesValues := make([]bool, 0)

	for g := 0; g < nGenes; g++ {
		if father.genes[g] == mother.genes[g] {
			commonGenes[g] = mother.genes[g]
		} else {
			unasignedGenesIndex = append(unasignedGenesIndex, g)
			unasignedGenesValues = append(unasignedGenesValues, mother.genes[g])
		}
	}

	firstChildGenes := getCompleteChildGenes(commonGenes, unasignedGenesIndex, unasignedGenesValues)
	secondChildGenes := getCompleteChildGenes(commonGenes, unasignedGenesIndex, unasignedGenesValues)

	firstChild = GenChromosomeFromGenes(firstChildGenes, distanceMatrix)
	secondChild = GenChromosomeFromGenes(secondChildGenes, distanceMatrix)

	return firstChild, secondChild
}
func UniformCrossover(father Chromosome, mother Chromosome, distanceMatrix [][]float32, m int) (firstChild, secondChild Chromosome) {
	nGenes := len(father.genes)
	firstChildGenes := make([]bool, nGenes, nGenes)
	secondChildGenes := make([]bool, nGenes, nGenes)

	for g := 0; g < nGenes; g++ {
		if father.genes[g] == mother.genes[g] {
			firstChildGenes[g] = mother.genes[g]
			secondChildGenes[g] = mother.genes[g]
		} else {
			conflictSolution := rand.Intn(4)
			switch conflictSolution {
			case 0:
				firstChildGenes[g] = mother.genes[g]
				secondChildGenes[g] = mother.genes[g]
			case 1:
				firstChildGenes[g] = father.genes[g]
				secondChildGenes[g] = father.genes[g]
			case 2:
				firstChildGenes[g] = mother.genes[g]
				secondChildGenes[g] = father.genes[g]
			case 3:
				firstChildGenes[g] = father.genes[g]
				secondChildGenes[g] = mother.genes[g]
			}
		}
	}

	repareGenes(firstChildGenes, distanceMatrix, m)
	repareGenes(secondChildGenes, distanceMatrix, m)

	firstChildValue := computeValue(distanceMatrix, firstChildGenes)
	secondChildValue := computeValue(distanceMatrix, secondChildGenes)

	return Chromosome{firstChildGenes, firstChildValue}, Chromosome{secondChildGenes, secondChildValue}

}

func Replace(current Poblation, selected Poblation) {
	sort.Sort(current)
	sort.Sort(selected)

	bestOfCurrent := current[len(current)-1]

	isStillSelected := false

	for i := 0; i < len(current) && !isStillSelected; i++ {
		isStillSelected = selected[i].Equals(bestOfCurrent)
	}

	if !isStillSelected {
		selected[0] = bestOfCurrent
	}
}

func Mutate(poblation Poblation, numOfMutations int, distanceMatrix [][]float32) {
	for i := 0; i < numOfMutations; i++ {
		indexToMutate := rand.Intn(len(poblation))
		poblation[indexToMutate].Mutate(distanceMatrix)
	}
}

func getCompleteChildGenes(commonGenes []bool, unasignedGenesIndex []int, unasignedGenesValues []bool) []bool {
	nGenes := len(commonGenes)
	toReturn := make([]bool, nGenes)

	for i := 0; i < nGenes; i++ {
		toReturn[i] = commonGenes[i]
	}

	shuffleGenes(unasignedGenesValues)

	for i, j := range unasignedGenesIndex {
		toReturn[j] = unasignedGenesValues[i]
	}

	return toReturn
}

func shuffleGenes(genes []bool) {
	rand.Shuffle(len(genes), func(i, j int) { genes[i], genes[j] = genes[j], genes[i] })
}

func getAWinner(currentPoblation Poblation) Chromosome {
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

func repareGenes(genes []bool, distanceMatrix [][]float32, m int) {
	currentlyTrue := countTrueGenes(genes)

	if currentlyTrue > m {
		discardGenes(genes, currentlyTrue-m, distanceMatrix)
	} else if currentlyTrue < m {
		addGenes(genes, m-currentlyTrue, distanceMatrix)
	}
}

func discardGenes(genes []bool, genesToDiscard int, distanceMatrix [][]float32) {

	mostContributors := getListOfMostContributors(genes, distanceMatrix)

	for i := 0; i < len(mostContributors) && genesToDiscard > 0; i++ {
		badContributor := mostContributors[i]
		genes[badContributor.index] = false
		genesToDiscard--
	}
}

func addGenes(genes []bool, genesToAdd int, distanceMatrix [][]float32) {
	mostPromissing := getListOfMostPromissing(genes, distanceMatrix)

	for i := 0; i < len(mostPromissing) && genesToAdd > 0; i++ {
		promissing := mostPromissing[i]
		genes[promissing.index] = true
		genesToAdd--
	}
}

type Contributor struct {
	index        int
	contribution float32
}

type Contributors []Contributor

func (cs Contributors) Len() int {
	return len(cs)
}

func (cs Contributors) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Contributors) Less(i, j int) bool {
	return cs[i].contribution < cs[j].contribution
}

func getListOfMostContributors(genes []bool, distanceMatrix [][]float32) Contributors {

	toReturn := make(Contributors, 0)

	for i, g := range genes {
		if g {
			contributor := Contributor{index: i, contribution: getContribution(genes, i, distanceMatrix)}
			toReturn = append(toReturn, contributor)
		}
	}

	sort.Sort(sort.Reverse(toReturn))

	return toReturn
}

func getContribution(genes []bool, contributor int, distanceMatrix [][]float32) float32 {
	var contribution float32 = 0
	for i, g := range genes {
		if g {
			contribution += distanceMatrix[i][contributor]
		}
	}
	return contribution
}

func getListOfMostPromissing(genes []bool, distanceMatrix [][]float32) Contributors {
	toReturn := make(Contributors, 0)
	for i, g := range genes {
		if !g {
			contributor := Contributor{index: i, contribution: getContribution(genes, i, distanceMatrix)}
			toReturn = append(toReturn, contributor)
		}
	}

	sort.Sort(sort.Reverse(toReturn))


	return toReturn
}

func countTrueGenes(genes []bool) int {
	toReturn := 0

	for _, g := range genes {
		if g {
			toReturn++
		}
	}
	return toReturn
}

func ChromosomeToSolution(chromosome Chromosome) []int {
	toReturn := make([]int, 0)

	for i, b := range chromosome.genes {
		if b {
			toReturn = append(toReturn, i)
		}
	}
	return toReturn
}
