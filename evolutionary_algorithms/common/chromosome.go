package common

import "math/rand"

type Chromosome struct {
	genes []bool
	value float32
}

func GenRandomChromosome(distanceMatrix [][]float32, n int, m int) Chromosome {

	genes := genRandomGenes(n, m)
	value := computeValue(distanceMatrix, genes)

	return Chromosome{genes, value}

}

func GenChromosomeFromGenes(genes []bool, distanceMatrix [][]float32) Chromosome {
	value := computeValue(distanceMatrix, genes)
	return Chromosome{genes, value}
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
