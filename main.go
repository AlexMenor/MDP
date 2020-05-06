package main

import (
	"MDP/evolutionary_algorithms/genetic_algorithm"
	"MDP/evolutionary_algorithms/memetic_algorithm"
	"MDP/greedy_algorithm"
	"MDP/local_search_algorithm"
	"MDP/problem_reader"
	"MDP/simulated_annealing_algorithm"
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	args := os.Args
	if len(args) > 1 && args[1] == "all" {
		runAllInstances()
	} else {
		runInteractive()
	}
}

func runInteractive() {

	instancesNames := getArrayOfInstancesNames()

	fmt.Println("Please choose a test!")
	printInstancesNames(instancesNames)

	choice := readChoice(len(instancesNames))

	file := path.Join("problem_instances", instancesNames[choice])

	n, m, distanceMatrix := problem_reader.ReadFile(file)

	printAlgorithmNames()

	choice = readChoice(9)

	var sol []int
	switch choice {
	case 0:
		sol = greedy_algorithm.Compute(n, m, distanceMatrix)
	case 1:
		sol, _ = local_search_algorithm.Compute(n, m, distanceMatrix)
	case 2:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Generational, genetic_algorithm.Positional)
	case 3:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Generational, genetic_algorithm.Uniform)
	case 4:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Stationary, genetic_algorithm.Positional)
	case 5:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Stationary, genetic_algorithm.Uniform)
	case 6:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.WholePoblation)
	case 7:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.OneRandom)
	case 8:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.BestOne)
	case 9:
		sol = simulated_annealing_algorithm.Compute(n,m, distanceMatrix)
	}

	fmt.Println(sol)
	fmt.Println(getDiversity(sol, distanceMatrix, m))

}

func runAllInstances() {

	const RESULTS_DIRECTORY = "results/"

	fileNames := []string{
		"greedy-results.csv",
		"local-search-results.csv",
		"genetic-generational-positional-results.csv",
		"genetic-generational-uniform-results.csv",
		"genetic-stationary-positional-results.csv",
		"genetic-stationary-uniform-results.csv",
		"memetic-whole-poblation.csv",
		"memetic-one-random.csv",
		"memetic-best-one.csv",

	}
	os.Mkdir(RESULTS_DIRECTORY, os.ModePerm)
	instancesNames := getArrayOfInstancesNames()

	for i, name := range fileNames {
		file, _ := os.Create(path.Join(RESULTS_DIRECTORY, name))
		for _, instance := range instancesNames {
			score, time := runInstance(instance, i)

			writeResultsCsv(instance, score, time, file)
		}

	}

}

func runInstance(instance string, algorithm int) (float32, int64) {
	file := path.Join("problem_instances", instance)

	n, m, distanceMatrix := problem_reader.ReadFile(file)

	var sol []int

	start := time.Now()
	switch algorithm {
	case 0:
		sol = greedy_algorithm.Compute(n, m, distanceMatrix)
	case 1:
		sol,_ = local_search_algorithm.Compute(n, m, distanceMatrix)
	case 2:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Generational, genetic_algorithm.Positional)
	case 3:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Generational, genetic_algorithm.Uniform)
	case 4:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Stationary, genetic_algorithm.Positional)
	case 5:
		sol = genetic_algorithm.Compute(distanceMatrix, n, m, 50, genetic_algorithm.Stationary, genetic_algorithm.Uniform)
	case 6:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.WholePoblation)
	case 7:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.OneRandom)
	case 8:
		sol = memetic_algorithm.Compute(distanceMatrix, n, m, 10, memetic_algorithm.BestOne)
	}

	end := time.Now()
	duration := end.Sub(start)

	return getDiversity(sol, distanceMatrix, m), duration.Microseconds()
}

func writeResultsCsv(instance string, score float32, time int64, file *os.File) {

	line := fmt.Sprintf("%s,%f,%d\n", instance, score, time)

	writter := bufio.NewWriter(file)
	writter.WriteString(line)
	writter.Flush()
}

func getArrayOfInstancesNames() []string {
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

func printInstancesNames(instancesNames []string) {
	for index, name := range instancesNames {
		fmt.Printf("%d. %s\n", index, name)
	}
}

func printAlgorithmNames() {
	fmt.Println("0. Greedy Algorithm")
	fmt.Println("1. Local Search Algorithm")
	fmt.Println("2. Genetic Generational Algorithm Positional Crossover")
	fmt.Println("3. Genetic Generational Algorithm Uniform Crossover")
	fmt.Println("4. Genetic Stationary Algorithm Positional Crossover")
	fmt.Println("5. Genetic Stationary Algorithm Uniform Crossover")
	fmt.Println("6. Memetic Algorithm (Whole Poblation)")
	fmt.Println("7. Memetic Algorithm (One Random)")
	fmt.Println("8. Memetic Algorithm (Best One)")
	fmt.Println("9. Simulated Annealing")
}

func getDiversity(selected []int, distanceMatrix [][]float32, m int) (diversity float32) {

	for i := 0; i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			diversity += distanceMatrix[selected[i]][selected[j]]
		}
	}

	return diversity
}

func readChoice(limit int) int {
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
