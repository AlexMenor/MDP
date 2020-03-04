package main

import (
	"MDP/greedy_algorithm"
	"MDP/local_search_algorithm"
	"MDP/problem_reader"
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const LOCAL_SEARCH_MAX_ITERATIONS = 100000

	instancesNames := getArrayOfInstancesNames()

	fmt.Println("Please choose a test!")
	printInstancesNames (instancesNames)

	choice := readChoice(len(instancesNames))

	file := path.Join("problem_instances", instancesNames[choice])

	n, m, distanceMatrix := problem_reader.ReadFile(file)

	printAlgorithmNames()

	choice = readChoice(2)

	var sol []int
	switch choice {
	case 0:
		sol = greedy_algorithm.Compute(n,m, distanceMatrix)
	case 1:
		sol = local_search_algorithm.Compute(n,m, distanceMatrix, LOCAL_SEARCH_MAX_ITERATIONS)
	}

	fmt.Println(sol)
	fmt.Println(getDiversity(sol, distanceMatrix, m))

}

func getArrayOfInstancesNames () []string {
	f, err := os.Open("./problem_instances")

	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	instancesNames := make([]string, 0)

	for _, file := range files {
		instancesNames = append(instancesNames, file.Name())
	}

	sort.Strings(instancesNames)

	return instancesNames
}

func printInstancesNames (instancesNames []string) {
	for index, name := range instancesNames {
		fmt.Printf("%d. %s\n", index, name)
	}
}

func printAlgorithmNames () {
	fmt.Println("0. Greedy Algorithm")
	fmt.Println("1. Local Search Algorithm")
}

func getDiversity(selected []int, distanceMatrix [][]float32, m int) (diversity float32) {

	for i := 0 ; i < m - 1 ; i++ {
		for j := i + 1 ; j < m ; j++ {
			diversity += distanceMatrix[selected[i]][selected[j]]
		}
	}

	return diversity
}

func readChoice (limit int) int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	stringArr := strings.Split(text, "\n")

	textChoice := stringArr[0]

	choice, err := strconv.Atoi(textChoice)

	if err != nil || choice < 0 || choice > limit {
		fmt.Print("That's not a valid choice!")
		os.Exit(1)
	}

	return choice
}