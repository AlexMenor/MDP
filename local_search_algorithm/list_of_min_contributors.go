package local_search_algorithm

import "sort"

type Contributor struct {
	key int
	contribution float32
}

func getListOfMinContributors (selected map[int]float32) []Contributor{

	sliceOfContributors := selectedSetToSlice(selected)

	sort.Slice(sliceOfContributors, func(i,j int) bool {
		return sliceOfContributors[i].contribution < sliceOfContributors[j].contribution
	})

	return sliceOfContributors
}

func selectedSetToSlice (selectedSet map[int]float32) []Contributor {
	sliceOfContributors := make([]Contributor, 0)

	for selected, contribution := range selectedSet {
		newContributor := Contributor{selected, contribution}
		sliceOfContributors = append(sliceOfContributors, newContributor)
	}

	return sliceOfContributors
}
