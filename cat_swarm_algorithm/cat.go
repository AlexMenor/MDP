package cat_swarm_algorithm

import "math/rand"

type Cat struct {
	x []bool
	fitness float32
	tracingMode bool
	velocityTrue []float32
	velocityFalse []float32
	distanceMatrix [][] float32
	n int
	m int
}

type Swarm []Cat

func GenerateSwarm(distanceMatrix [][]float32, n, m, nCats int) Swarm {
	swarm := make(Swarm, 0,0)

	for i := 0 ; i < nCats ; i++ {
		cat := genRandomCat(distanceMatrix,n,m)
		swarm = append(swarm, cat)
	}

	return swarm
}

func (cat * Cat) RateCat() {

	newFitness := computeValue(cat.distanceMatrix, cat.x)

	cat.fitness = newFitness
}

func (swarm Swarm) GetBestGenes () ([]bool, float32) {

	bestFoundGenes := make([]bool, swarm[0].n)
	var bestFoundFitness float32 = 0

	for i := 0 ; i < len(swarm) ; i++ {
		cat := swarm[i]
		if cat.fitness > bestFoundFitness {
			bestFoundFitness = cat.fitness
			copy(bestFoundGenes, cat.x)
		}
	}

	return bestFoundGenes, bestFoundFitness
}

func (swarm * Swarm) SetMixtureRatio (mixtureRatio float32) {

	swarm.setCatsTracingToFalse()

	catsToSetTracing := int(mixtureRatio * float32(len(*swarm)))

	setOfIndexes := make(map[int]bool)

	for catsToSetTracing > len(setOfIndexes) {
		random := rand.Intn(len(*swarm))
		setOfIndexes[random] = true
	}

	for i := range setOfIndexes {
		(*swarm)[i].tracingMode = true
	}
}

func (swarm * Swarm) setCatsTracingToFalse () {

	for i := 0 ; i < len(*swarm) ; i++ {
		cat := (*swarm)[i]

		cat.tracingMode = false
	}
}

func computeValue(distanceMatrix [][]float32, genes []bool) float32 {
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

func genRandomCat (distanceMatrix[][]float32, n,m int) Cat {
	x := genRandomGenes(n,m)
	velocityTrue := genVelocity(n)
	velocityFalse := genVelocity(n)

	cat := Cat {x, 0, false, velocityTrue, velocityFalse, distanceMatrix, n,m}

	cat.RateCat()

	return cat
}

func genVelocity(n int) []float32 {
	velocity := make([]float32, n)

	return velocity
}

func genRandomGenes(n int, m int) []bool {
	genes := make([]bool, n, n)

	i := 0
	for i < m {
		randomPos := rand.Intn(n)
		if !genes[randomPos] {
			genes[randomPos] = true
			i++
		}
	}

	return genes
}
