package evolutionary_algorithms

import (
	"MDP/evolutionary_algorithms/common"
	"math"
)

type GeneticType string
type CrossoverType string

const (
	Generational GeneticType   = "GENERATIONAL"
	Stationary   GeneticType   = "STATIONARY"
	Uniform      CrossoverType = "UNIFORM"
	Positional   CrossoverType = "POSITIONAL"
)

func Compute(distanceMatrix [][]float32, n, m, tam int, geneticType GeneticType, crossoverType CrossoverType) common.Chromosome {
	if geneticType == Generational {
		return computeGenerational(distanceMatrix, n, m, tam, crossoverType)
	}
}

func computeGenerational(distanceMatrix [][]float32, n, m, tam int, crossoverType CrossoverType) common.Chromosome {
	crossoverProbability := 0.7
	MAX_TARGET_EVALUATIONS := 100000

	var numOfCrossoversPerIteration int = int(math.Floor(crossoverProbability * (float64(tam) / 2)))

	var iterationsLeft int = MAX_TARGET_EVALUATIONS - tam

	MAX_ITERATIONS := iterationsLeft / numOfCrossoversPerIteration

	poblation := common.GenRandomPoblation(distanceMatrix, n, m, tam)

	for i := 0; i < MAX_ITERATIONS; i++ {
		selected := common.GetSelectedFromPoblation(poblation, tam)
		for j := 0 ; j < numOfCrossoversPerIteration * 2 ; j+=2 {
			father := selected[j]
			mother := selected[j+1]

			var firstChild common.Chromosome
			var secondChild common.Chromosome

		}
	}

}
