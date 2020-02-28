package problem_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile (path string) (nCount, mCount int, distanceMatrix [][] float32){
	file, err := os.Open(path)

	checkForErrors(err)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()

	nCount = extractN(line)
	mCount = extractM(line)

	distanceMatrix = make([][]float32, nCount)
	for i := range distanceMatrix {
		distanceMatrix[i] = make([]float32, nCount)
	}

	for scanner.Scan() {
		line = scanner.Text()
		writeDistance(line, distanceMatrix)
	}


	return nCount, mCount, distanceMatrix
}

func writeDistance (line string, distanceMatrix [][]float32){
	arr := strings.Split(line, " ")

	i, err := strconv.Atoi(arr[0])
	checkForErrors(err)
	j, err := strconv.Atoi(arr[1])
	checkForErrors(err)
	distance, err := strconv.ParseFloat(arr[2], 32)
	checkForErrors(err)

	(distanceMatrix)[i][j] = float32(distance)
	(distanceMatrix)[j][i] = float32(distance)
}

func checkForErrors(err error){
	if err != nil {
		panic(err)
	}
}

func extractN (line string) int{
	arr := strings.Split(line, " ")

	n, err := strconv.Atoi(arr[0])

	checkForErrors(err)

	return n
}

func extractM (line string) int{
	arr := strings.Split(line, " ")

	m, err := strconv.Atoi(arr[1])

	checkForErrors(err)

	return m
}
