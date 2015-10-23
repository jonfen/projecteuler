package main

import (
	"testing"
)

var testSet = []struct {
	dataGrid string
	result   int
}{
	{`data_grid.txt`, 70600674},
}

func TestFunc(t *testing.T) {
	for _, tt := range testSet {
		if x := searchDataGrid(loadDataGrid(tt.dataGrid)); x != tt.result {
			t.Errorf("The largest product in the grid should be %v, not %v", tt.result, x)
		}
	}
}

func BenchmarkFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		searchDataGrid(loadDataGrid(`data_grid.txt`))
	}
}
