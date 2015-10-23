package main

import (
	"bufio"
	"flag"
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
	"io"
	"os"
	"strconv"
	"strings"
)

var count int

var answer Answer

type dataGrid [20][20]int

type Factor struct {
	x, y int
}

type Answer struct {
	Factors [4]Factor
}

func (a *Answer) product(dg dataGrid) int {
	return dg[a.Factors[0].x][a.Factors[0].y] * dg[a.Factors[1].x][a.Factors[1].y] * dg[a.Factors[2].x][a.Factors[2].y] * dg[a.Factors[3].x][a.Factors[3].y]
}

func (a *Answer) print(dg dataGrid) {
	fmt.Print("\n", dg[a.Factors[0].x][a.Factors[0].y], "*", dg[a.Factors[1].x][a.Factors[1].y], "*", dg[a.Factors[2].x][a.Factors[2].y], "*", dg[a.Factors[3].x][a.Factors[3].y], "=", a.product(dg))
}

func printDataGrid(dg dataGrid, answer Answer) {
	for x, row := range dg {
		for y, cell := range row {
			factor := Factor{x: x, y: y}
			if factor == answer.Factors[0] || factor == answer.Factors[1] || factor == answer.Factors[2] || factor == answer.Factors[3] {
				ct.Foreground(ct.Red, false)
			}

			if cell > 9 {
				fmt.Printf("%v ", cell)
			} else {
				fmt.Printf("0%v ", cell)
			}
			ct.Foreground(ct.White, false)

		}
		fmt.Println()
	}
}

func init() {
	flag.IntVar(&count, "count", 13, "consecutive number count")
}

// Error checking helper function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadDataGrid(filename string) dataGrid {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	dg := dataGrid{}
	lineNumber := 0
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	for err == nil {
		lineArray := strings.Split(line, " ")
		for i, v := range lineArray {
			dg[lineNumber][i], err = strconv.Atoi(strings.Replace(v, "\n", "", 1))
			check(err)
		}
		line, err = r.ReadString('\n')
		lineNumber++
	}
	if err != io.EOF {
		fmt.Println(err)
		return dg
	}
	return dg
}

func searchDataGrid(dg dataGrid) int {
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {

			// Check horizontal left
			if x > 3 {
				factors := Answer{
					Factors: [4]Factor{
						Factor{y: y, x: x},
						Factor{y: y, x: x - 1},
						Factor{y: y, x: x - 2},
						Factor{y: y, x: x - 3}},
				}

				if factors.product(dg) > answer.product(dg) {
					answer = factors
				}
			}

			// Check vertical up
			if y > 3 {
				factors := Answer{
					Factors: [4]Factor{
						Factor{y: y, x: x},
						Factor{y: y - 1, x: x},
						Factor{y: y - 2, x: x},
						Factor{y: y - 3, x: x}},
				}

				if factors.product(dg) > answer.product(dg) {
					answer = factors
				}
			}

			// Check diagonal up left
			if y > 3 && x > 3 {
				factors := Answer{
					Factors: [4]Factor{
						Factor{y: y, x: x},
						Factor{y: y - 1, x: x - 1},
						Factor{y: y - 2, x: x - 2},
						Factor{y: y - 3, x: x - 3}},
				}

				if factors.product(dg) > answer.product(dg) {
					answer = factors
				}
			}

			// Check diagonal up right
			if y > 3 && x < 17 {
				factors := Answer{
					Factors: [4]Factor{
						Factor{y: y, x: x},
						Factor{y: y - 1, x: x + 1},
						Factor{y: y - 2, x: x + 2},
						Factor{y: y - 3, x: x + 3}},
				}

				if factors.product(dg) > answer.product(dg) {
					answer = factors
				}
			}
		}
	}
	return answer.product(dg)
}

// https://projecteuler.net/problem=11
func main() {
	flag.Parse()

	dg := loadDataGrid(`data_grid.txt`)

	fmt.Println("In the 20×20 grid below, four numbers along a diagonal line have been marked in red.\n")
	factors := Answer{
		Factors: [4]Factor{
			Factor{y: 8, x: 6},
			Factor{y: 9, x: 7},
			Factor{y: 10, x: 8},
			Factor{y: 11, x: 9}},
	}

	printDataGrid(dg, factors)
	fmt.Println("\nThe product of these numbers is 26 × 63 × 78 × 14 = 1788696.\n")

	searchDataGrid(dg)

	printDataGrid(dg, answer)
	fmt.Printf("\nThe greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid is %v.\n\n", answer.product(dg))
}
