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
	var bestChromosome common.Chromosome
	//if geneticType == Generational {
		bestChromosome = computeGenerational(distanceMatrix, n, m, tam, crossoverType)
	//}
	return common.ChromosomeToSolution(bestChromosome)
}

func computeGenerational(distanceMatrix [][]float32, n, m, tam int, crossoverType CrossoverType) common.Chromosome {
	crossoverProbability := 0.7
	mutationProbabilityPerGene := 0.001
	MAX_TARGET_EVALUATIONS := 100000

	var numOfCrossoversPerIteration int = int(math.Floor(crossoverProbability * (float64(tam) / 2)))
	var numOfMutationsPerIteration int = int(math.Floor(mutationProbabilityPerGene * float64(n * tam)))

	var evaluationsLeft int = MAX_TARGET_EVALUATIONS - tam

	MAX_ITERATIONS := evaluationsLeft / (numOfCrossoversPerIteration + numOfMutationsPerIteration)

	poblation := common.GenRandomPoblation(distanceMatrix, n, m, tam)

	for i := 0; i < MAX_ITERATIONS; i++ {
		selected := common.GetSelectedFromPoblation(poblation)
		for j := 0; j < numOfCrossoversPerIteration*2; j += 2 {
			father := selected[j]
			mother := selected[j+1]

			var firstChild common.Chromosome
			var secondChild common.Chromosome

			if crossoverType == Positional {
				firstChild, secondChild = common.PositionalCrossOver(father, mother, distanceMatrix)
			} else {
				firstChild, secondChild = common.UniformCrossOver(father, mother, distanceMatrix)
			}

			selected[j] = firstChild
			selected[j+1] = secondChild
		}

		common.Mutate(selected, numOfMutationsPerIteration, distanceMatrix)

		common.Replace(poblation, selected)
	}

	sort.Sort(poblation)

	return poblation[len(poblation)- 1]

}
