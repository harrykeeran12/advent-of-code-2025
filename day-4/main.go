package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Read file
// Convert a file into a two dimensional array of characters
func arrayInit(filename string) [][]string {
	f, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	file := string(f)

	fileArray := strings.Split(file, "\n")

	var arr [][]string
	for lineIndex := range fileArray {
		line := fileArray[lineIndex]
		if line != "" {
			// Turn this into an array
			charSlice := strings.Split(line, "")
			arr = append(arr, charSlice)

		}
	}
	return arr

}

// This function checks if you can access a specific element by checking the ones around it.
func canAccess(arr [][]string, row int, column int) bool {
	rowNum := len(arr)
	columnNum := len(arr[0])
	// fmt.Println("Max number of rows is", rowNum, "Max columns is", columnNum)
	// if (row > 0 || column > 0) && (row < rowNum - 1 || column < columnNum - 1){
	// This means not the borders.
	element := arr[row][column]
	count := 0
	if element == "@" {
		for i := range 3 {
			for j := range 3 {
				currRow := (row - 1) + i
				currColumn := (column - 1) + j
				// fmt.Println(currRow, currColumn)
				if !(currRow == row && currColumn == column) {
					// if (currRow >= 0 || currColumn >= 0) && (currRow <= rowNum - 1 || currColumn <= columnNum - 1){
					// }
					if (currRow >= 0 && currRow <= rowNum-1) && (currColumn >= 0 && currColumn <= columnNum-1) {
						// fmt.Println(currRow, currColumn)
						if arr[currRow][currColumn] == "@" {
							count += 1
						}

					}
				}
			}
		}
		if count < 4 {
			// fmt.Println(count)
			return true
		} else {
			return false
		}
	}
	return false
}

func forkLiftAccess(filename string) int {
	arr := arrayInit(filename)
	rowNum := len(arr)
	columnNum := len(arr[0])
	for line := range arr {
		fmt.Println(arr[line])
	}
	fmt.Println("----")
	count := 0
	for i := range rowNum {
		for j := range columnNum {
			if arr[i][j] == "@" {
				access := canAccess(arr, i, j)
				if access {
					// arr[i][j] = "x"
					count++
				}
			}
			// fmt.Printf("Can access arr[%d][%d] = %t\n", i, j, access)
		}
	}
	// for line := range arr{
	// 	fmt.Println(arr[line])
	// }
	fmt.Println("Number of accessable rolls =", count)
	return count

}

func forkliftaccessTwo(parr *[][]string)(int, [][]int){
	// arr := arrayInit(filename)
	arr := (*parr)
	var indexes [][]int
	rowNum := len(arr)
	columnNum := len(arr[0])
	for line := range arr {
		fmt.Println(arr[line])
	}
	fmt.Println("----")
	count := 0
	for i := range rowNum {
		for j := range columnNum {
			if arr[i][j] == "@" {
				access := canAccess(arr, i, j)
				if access {
					// arr[i][j] = "x"
					temp := []int{i, j}
					indexes = append(indexes,temp)
					count++
				}
			}
			// fmt.Printf("Can access arr[%d][%d] = %t\n", i, j, access)
		}
	}
	// for line := range arr{
	// 	fmt.Println(arr[line])
	// }
	fmt.Println("Number of accessable rolls =", count)
	return count, indexes
}

func applyIndexes(parr *[][]string, indexes [][]int) ([][]string){
	arr := (*parr)
	for i := range indexes {
		numbers := indexes[i]
		arr[numbers[0]][numbers[1]] = "."
		}
	return arr
}
// Repeatedly call the forkliftaccess
func part2(filename string){
	sum := 0
	arr := arrayInit(filename)
	var prevCount int
	count := -1
	for count != 1 && count != 0{
		count, indexes := forkliftaccessTwo(&arr)
		if count == 0{
			break
		}
		arr = applyIndexes(&arr, indexes)
		sum += count
		prevCount = count
	}
	fmt.Println("Sum is ", sum)
	fmt.Println(prevCount)
	// fmt.Println(count, indexes)

	// fmt.Println(forkliftaccessTwo(&arr))



}

func main() {
	fmt.Printf("Hello World.")
}
