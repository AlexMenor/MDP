package cat_swarm_algorithm

import (
	"MDP/local_search_algorithm"
	"math"
	"math/rand"
)

func ComputeMemetic(distanceMatrix [][]float32, n, m int) []int {

	N_CATS := 30

	// Cats in Tracing Mode
	var MIXTURE_RATE float32 = 0.2

	// Number of candidates in Seeking Mode
	SMP := 3

	// Count current position as candidate in Seeking Mode
	SPC := true

	// Max number of dimensions able to mutate in Seeking Mode
	CDC := 3

	// Probability of mutation in Seeking Mode
	PMO := 0.2

	// Inertia Weight in Tracing Mode
	W := 0.1

	// Constant in Tracing Mode
	C1 := 10

	// Max Evaluations
	MAX_EVALUATIONS := 100000

	// Calculate max iterations
	MAX_ITERATIONS := float32(MAX_EVALUATIONS) / (float32(N_CATS) + (1-MIXTURE_RATE)*float32(N_CATS)*float32(SMP) + MIXTURE_RATE*float32(N_CATS))

	swarm := GenerateSwarm(distanceMatrix, n, m, N_CATS)

	var bestGenes []bool

	for i := 0; i < int(MAX_ITERATIONS); i++ {
		bestGenes, _ = swarm.GetBestGenes()

		swarm.SetMixtureRatio(MIXTURE_RATE)

		for j := 0; j < N_CATS; j++ {
			cat := swarm[j]
			if cat.tracingMode {
				cat.ApplyTracing(bestGenes, float32(C1), float32(W))

			} else {
				cat.ApplySeeking(SMP, CDC, float32(PMO), SPC)
			}
		}

		if i % 10 == 0 {
			solution := genesToSolution(bestGenes)
			localSearchResult, _ := local_search_algorithm.ComputeForMemeticAlgorithm(n,m, distanceMatrix,solution)
			genes := genesFromSolution(n, localSearchResult)
			swarm[0].x = genes
			swarm[0].RateCat()
		}
	}

	bestGenes, _ = swarm.GetBestGenes()

	return genesToSolution(bestGenes)
}

func Compute(distanceMatrix [][]float32, n, m int) []int {

	N_CATS := 30

	// Cats in Tracing Mode
	var MIXTURE_RATE float32 = 0.2

	// Number of candidates in Seeking Mode
	SMP := 3

	// Count current position as candidate in Seeking Mode
	SPC := true

	// Max number of dimensions able to mutate in Seeking Mode
	CDC := 3

	// Probability of mutation in Seeking Mode
	PMO := 0.2

	// Inertia Weight in Tracing Mode
	W := 0.1

	// Constant in Tracing Mode
	C1 := 10

	// Max Evaluations
	MAX_EVALUATIONS := 100000

	// Calculate max iterations
	MAX_ITERATIONS := float32(MAX_EVALUATIONS) / (float32(N_CATS) + (1-MIXTURE_RATE)*float32(N_CATS)*float32(SMP) + MIXTURE_RATE*float32(N_CATS))

	swarm := GenerateSwarm(distanceMatrix, n, m, N_CATS)

	var bestGenes []bool

	for i := 0; i < int(MAX_ITERATIONS); i++ {
		bestGenes, _ = swarm.GetBestGenes()

		swarm.SetMixtureRatio(MIXTURE_RATE)

		for j := 0; j < N_CATS; j++ {
			cat := swarm[j]
			if cat.tracingMode {
				cat.ApplyTracing(bestGenes, float32(C1), float32(W))

			} else {
				cat.ApplySeeking(SMP, CDC, float32(PMO), SPC)
			}
		}
	}

	bestGenes, _ = swarm.GetBestGenes()

	return genesToSolution(bestGenes)
}

func (cat * Cat) ApplyTracing(bestGenes [] bool, c1 float32, w float32) {
	cat.UpdateVelocities(bestGenes, c1, w)

	indexesToMutateToOne := make([]int,0)
	indexesToMutateToZero := make([]int,0)

	for i := 0 ; i < cat.n ; i++{
		var v float32
		if cat.x[i] {
			v = cat.velocityFalse[i]
		} else {
			v = cat.velocityTrue[i]
		}

		t := 1 / (1 + math.Exp(float64(-v)))

		if rand.Float64() < t {
			if cat.x[i]	{
				indexesToMutateToZero = append(indexesToMutateToZero, i)
			} else {
				indexesToMutateToOne = append(indexesToMutateToOne, i)
			}
		}
	}

	rand.Shuffle(len(indexesToMutateToZero), func(i, j int) { indexesToMutateToZero[i], indexesToMutateToZero[j] = indexesToMutateToZero[j], indexesToMutateToZero[i] })
	rand.Shuffle(len(indexesToMutateToOne), func(i, j int) { indexesToMutateToOne[i], indexesToMutateToOne[j] = indexesToMutateToOne[j], indexesToMutateToOne[i] })

	minLength := len(indexesToMutateToZero)
	if len(indexesToMutateToZero) > len(indexesToMutateToOne) {
		minLength = len(indexesToMutateToOne)
	}

	for i:= 0 ; i < minLength ; i++ {
		indexToZero := indexesToMutateToZero[i]
		indexToOne := indexesToMutateToOne[i]

		cat.x[indexToZero] = false
		cat.x[indexToOne] = true
	}

	cat.RateCat()
}

func (cat * Cat) UpdateVelocities(bestGenes[] bool, c1 float32, w float32) {
	for i := 0 ; i < cat.n ; i++ {
		dOne := -rand.Float32() * c1
		dZero := rand.Float32() * c1

		if bestGenes[i] {
			dOne = -dOne
			dZero = -dZero
		}

		cat.velocityTrue[i] = cat.velocityTrue[i] * w + dOne
		cat.velocityFalse[i] = cat.velocityFalse[i] * w + dZero
	}
}

func (cat *Cat) ApplySeeking(numberOfCandidates int, numberOfMutations int, probabilityOfMutation float32, preserveOriginal bool) {
	candidates := make([]Candidate, 0)
	copyOfCatGenes := make([]bool,  cat.n, cat.n)

	minFitness := cat.fitness

	copy(copyOfCatGenes, cat.x)

	for i := 0; i < numberOfCandidates; i++ {
		newCandidate := Candidate{copyOfCatGenes, cat.fitness, 0.0}
		if !preserveOriginal || i != 0 {
			newCandidate.Mutate(numberOfMutations, probabilityOfMutation, cat.distanceMatrix)
		}
		candidates = append(candidates, newCandidate)

		if newCandidate.fitness < minFitness {
			minFitness = newCandidate.fitness
		}
	}

	for i := 0; i < numberOfCandidates; i++ {

		candidates[i].probabilityOfBeingChosen = 100 * ((candidates[i].fitness - minFitness) / candidates[i].fitness) + 0.5
	}

	chosenCandidate := rouletteSelect(candidates)

	cat.x = chosenCandidate.x
	cat.fitness = chosenCandidate.fitness
}

func (candidate *Candidate) Mutate(numOfMutations int, probabilityOfMutation float32, distanceMatrix [][]float32) {
	isMutated := false

	for i := 0; i < numOfMutations; i++ {
		if rand.Float32() < probabilityOfMutation {
			isMutated = true

			genToMutate := rand.Intn(len(candidate.x))
			current := candidate.x[genToMutate]
			newValue := !current

			otherGenToMutate := rand.Intn(len(candidate.x))
			for candidate.x[otherGenToMutate] == current {
				otherGenToMutate = rand.Intn(len(candidate.x))
			}

			candidate.x[genToMutate], candidate.x[otherGenToMutate] = newValue, current
		}
	}

	if isMutated {
		candidate.fitness = computeGenes(distanceMatrix, candidate.x)
	}
}

type Candidate struct {
	x                        []bool
	fitness                  float32
	probabilityOfBeingChosen float32
}

func computeGenes(distanceMatrix [][]float32, genes []bool) float32 {
	n := len(distanceMatrix)
	var toReturn float32 = 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n && genes[i]; j++ {
			if genes[j] {
				toReturn += distanceMatrix[i][j]
			}
		}
	}
	return toReturn
}

func rouletteSelect(candidates []Candidate) Candidate {
	var sum float32 = 0.0
	for _, candidate := range candidates {
		sum += candidate.probabilityOfBeingChosen
	}
	value := rand.Float32() * sum

	for _, candidate := range candidates {
		value -= candidate.probabilityOfBeingChosen
		if value <= 0 {
			return candidate
		}
	}
	return candidates[len(candidates)-1]
}
func genesToSolution(genes []bool) []int {
	toReturn := make([]int, 0)

	for i, b := range genes {
		if b {
			toReturn = append(toReturn, i)
		}
	}
	return toReturn
}

func genesFromSolution(n int, solution []int) []bool {
	genes := make([]bool, n, n)

	for _, s := range solution {
		genes[s] = true
	}

	return genes
}
