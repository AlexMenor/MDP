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
	"time"
)

func main() {

	args := os.Args
	if len(args) > 1 {
		if args[1] == "all" {
			runAllInstances()
		} else {
			runInteractive()
		}
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

	choice = readChoice(2)

	var sol []int
	switch choice {
	case 0:
		sol = greedy_algorithm.Compute(n, m, distanceMatrix)
	case 1:
		sol = local_search_algorithm.Compute(n, m, distanceMatrix)
	}

	fmt.Println(sol)
	fmt.Println(getDiversity(sol, distanceMatrix, m))

}

func runAllInstances() {

	const RESULTS_DIRECTORY = "results/"

	const GREEDY_FILE_NAME = "greedy-results.csv"
	const LOCAL_SEARCH_FILE_NAME = "local-search-results.csv"

	os.Mkdir(RESULTS_DIRECTORY, os.ModePerm)

	greedyFile, _ := os.Create(path.Join(RESULTS_DIRECTORY, GREEDY_FILE_NAME))
	localSearchFile, _ := os.Create(path.Join(RESULTS_DIRECTORY, LOCAL_SEARCH_FILE_NAME))

	instancesNames := getArrayOfInstancesNames()

	for _, instance := range instancesNames {
		scoreGreedy, scoreLocal, timeGreedy, timeLocal := runInstance(instance)

		writeResultsCsv(instance, scoreGreedy, timeGreedy, greedyFile)
		writeResultsCsv(instance, scoreLocal, timeLocal, localSearchFile)
	}
}

func runInstance(instance string) (float32, float32, int64, int64) {
	file := path.Join("problem_instances", instance)

	n, m, distanceMatrix := problem_reader.ReadFile(file)

	start := time.Now()
	solGreedy := greedy_algorithm.Compute(n, m, distanceMatrix)
	end := time.Now()
	durationGreedy := end.Sub(start)

	start = time.Now()
	solLocal := local_search_algorithm.Compute(n, m, distanceMatrix)
	end = time.Now()
	durationLocal := end.Sub(start)

	return getDiversity(solGreedy, distanceMatrix, m), getDiversity(solLocal, distanceMatrix, m),
	durationGreedy.Microseconds(), durationLocal.Microseconds()
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
