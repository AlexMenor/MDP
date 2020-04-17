package memetic_algorithm

import (
	"MDP/evolutionary_algorithms/common"
	"MDP/local_search_algorithm"
	"math"
	"math/rand"
	"sort"
)

type MemeticType string

const (
	WholePoblation MemeticType = "WHOLE_POBLATION"
	OneRandom      MemeticType = "ONE_RANDOM"
	BestOne        MemeticType = "BEST_ONE"
)

func Compute(distanceMatrix [][]float32, n, m, tam int, memeticType MemeticType) []int {
	crossoverProbability := 0.7
	mutationProbabilityPerGene := 0.001
	MAX_TARGET_EVALUATIONS := 100000

	var numOfCrossoversPerIteration int = int(math.Floor(crossoverProbability * (float64(tam) / 2)))
	var numOfMutationsPerIteration int = int(math.Floor(mutationProbabilityPerGene * float64(n*tam)))

	poblation := common.GenRandomPoblation(distanceMatrix, n, m, tam)

	currentEvaluations := 0
	hybridCountDown := 10

	for currentEvaluations < MAX_TARGET_EVALUATIONS {
		selected := common.GetSelectedFromPoblation(poblation, tam)
		for j := 0; j < numOfCrossoversPerIteration*2; j += 2 {
			father := selected[j]
			mother := selected[j+1]

			var firstChild common.Chromosome
			var secondChild common.Chromosome

			firstChild, secondChild = common.UniformCrossover(father, mother, distanceMatrix, m)

			selected[j] = firstChild
			selected[j+1] = secondChild
		}
		currentEvaluations += numOfCrossoversPerIteration * 2

		common.Mutate(selected, numOfMutationsPerIteration, distanceMatrix)

		currentEvaluations += numOfMutationsPerIteration

		poblation = common.KeepBest(poblation, selected)

		hybridCountDown--

		if hybridCountDown == 0 {
			if memeticType == WholePoblation {
				for i := 0; i < len(poblation); i++ {
					currentSolution := common.ChromosomeToSolution(poblation[i])
					localSolution, evaluations := local_search_algorithm.ComputeForMemeticAlgorithm(n, m, distanceMatrix, currentSolution)
					poblation[i] = common.GenChromosomeFromSolution(localSolution, distanceMatrix)
					currentEvaluations += evaluations
				}
			} else if memeticType == OneRandom {
				randomIndex := rand.Intn(tam)
				randomChromosome := poblation[randomIndex]
				currentSolution := common.ChromosomeToSolution(randomChromosome)
				localSolution, evaluations := local_search_algorithm.ComputeForMemeticAlgorithm(n, m, distanceMatrix, currentSolution)
				poblation[randomIndex] = common.GenChromosomeFromSolution(localSolution, distanceMatrix)
				currentEvaluations += evaluations
			} else {
				bestChromosome := poblation[tam-1]
				currentSolution := common.ChromosomeToSolution(bestChromosome)
				localSolution, evaluations := local_search_algorithm.ComputeForMemeticAlgorithm(n, m, distanceMatrix, currentSolution)
				poblation[tam-1] = common.GenChromosomeFromSolution(localSolution, distanceMatrix)
				currentEvaluations += evaluations
			}
			hybridCountDown = 10
		}
	}

	sort.Sort(poblation)
	bestChromosome := poblation[tam-1]

	return common.ChromosomeToSolution(bestChromosome)
}
