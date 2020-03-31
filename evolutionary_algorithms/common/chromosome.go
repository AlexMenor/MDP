package common

import "math/rand"

type Chromosome struct {
	genes []bool
	value float32
}

func (c Chromosome) Equals(other Chromosome) bool {
	if len(c.genes) != len(other.genes) {
		return false
	}
	toReturn := true

	for i := 0; i < len(c.genes) && toReturn; i++ {
		toReturn = c.genes[i] == other.genes[i]
	}

	return toReturn
}

func (c *Chromosome) Mutate(distanceMatrix [][]float32) {
	genToMutate := rand.Intn(len(c.genes))
	current := c.genes[genToMutate]
	newValue := !current

	otherGenToMutate := rand.Intn(len(c.genes))
	for otherGenToMutate == genToMutate || c.genes[otherGenToMutate] == current {
		otherGenToMutate = rand.Intn(len(c.genes))
	}

	c.genes[genToMutate] = newValue
	c.genes[otherGenToMutate] = current

	c.value = computeValue(distanceMatrix, c.genes)
}

type Poblation []Chromosome

func (p Poblation) Len() int {
	return len(p)
}

func (p Poblation) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Poblation) Less(i, j int) bool {
	return p[i].value < p[j].value
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
