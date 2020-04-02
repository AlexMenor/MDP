package genetic_algorithm

import (
	"MDP/evolutionary_algorithms/common"
	"math"
	"sort"
)

type GeneticType string
type CrossoverType string

const (
	Generational GeneticType   = "GENERATIONAL"
	Stationary   GeneticType   = "STATIONARY"
	Uniform      CrossoverType = "UNIFORM"
	Positional   CrossoverType = "POSITIONAL"
)

func Compute(distanceMatrix [][]float32, n, m, tam int, geneticType GeneticType, crossoverType CrossoverType) []int {
	var poblation common.Poblation
	if geneticType == Generational {
		poblation = computeGenerational(distanceMatrix, n, m, tam, crossoverType)
	} else {
		poblation = computeStationary(distanceMatrix, n, m, tam, crossoverType)
	}

	sort.Sort(poblation)
	bestChromosome := poblation[len(poblation)-1]

	return common.ChromosomeToSolution(bestChromosome)

}

func computeGenerational(distanceMatrix [][]float32, n, m, tam int, crossoverType CrossoverType) common.Poblation {
	crossoverProbability := 0.7
	mutationProbabilityPerGene := 0.001
	MAX_TARGET_EVALUATIONS := 100000

	var numOfCrossoversPerIteration int = int(math.Floor(crossoverProbability * (float64(tam) / 2)))
	var numOfMutationsPerIteration int = int(math.Floor(mutationProbabilityPerGene * float64(n*tam)))

	var evaluationsLeft int = MAX_TARGET_EVALUATIONS - tam

	MAX_ITERATIONS := evaluationsLeft / (numOfCrossoversPerIteration*2 + numOfMutationsPerIteration)

	poblation := common.GenRandomPoblation(distanceMatrix, n, m, tam)

	for i := 0; i < MAX_ITERATIONS; i++ {
		selected := common.GetSelectedFromPoblation(poblation, tam)
		for j := 0; j < numOfCrossoversPerIteration*2; j += 2 {
			father := selected[j]
			mother := selected[j+1]

			var firstChild common.Chromosome
			var secondChild common.Chromosome

			if crossoverType == Positional {
				firstChild, secondChild = common.PositionalCrossover(father, mother, distanceMatrix)
			} else {
				firstChild, secondChild = common.UniformCrossover(father, mother, distanceMatrix, m)
			}

			selected[j] = firstChild
			selected[j+1] = secondChild
		}

		common.Mutate(selected, numOfMutationsPerIteration, distanceMatrix)

		poblation = common.KeepBest(poblation, selected)
	}

	return poblation
}

func computeStationary(distanceMatrix [][]float32, n, m, tam int, crossoverType CrossoverType) common.Poblation {
	mutationProbabilityPerGene := 0.001
	MAX_TARGET_EVALUATIONS := 100000

	numOfParentsPerIteration := 2
	var numOfMutationsPerIteration int = int(math.Floor(mutationProbabilityPerGene * float64(n*numOfParentsPerIteration)))

	var evaluationsLeft int = MAX_TARGET_EVALUATIONS - tam

	MAX_ITERATIONS := evaluationsLeft / (numOfParentsPerIteration + numOfMutationsPerIteration)

	poblation := common.GenRandomPoblation(distanceMatrix, n, m, tam)

	for i := 0; i < MAX_ITERATIONS; i++ {
		selected := common.GetSelectedFromPoblation(poblation, numOfParentsPerIteration)

		father := selected[0]
		mother := selected[1]

		var firstChild common.Chromosome
		var secondChild common.Chromosome

		if crossoverType == Positional {
			firstChild, secondChild = common.PositionalCrossover(father, mother, distanceMatrix)
		} else {
			firstChild, secondChild = common.UniformCrossover(father, mother, distanceMatrix, m)
		}

		selected[0] = firstChild
		selected[1] = secondChild
		common.Mutate(selected, numOfMutationsPerIteration, distanceMatrix)

		sort.Sort(poblation)

		poblation[0] = selected[0]
		poblation[1] = selected[1]
	}

	return poblation
}
