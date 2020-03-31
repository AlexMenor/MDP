package common

import "testing"

func TestGenRandomGenes(t *testing.T) {
	n := 500
	m := 50
	got := genRandomGenes(n,m)

	if len(got) != n {
		t.Errorf("Genes tam was not %d as expected", n)
	}

	counter := 0

	for _, gen := range got {
		if gen {
			counter++
		}
	}

	if counter != m {
		t.Errorf("There was no %d positive genes as expected", m)
	}
}

