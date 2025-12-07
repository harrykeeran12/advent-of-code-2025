package day5

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	// Take the string, and process it into two arrays, the list of ranges and the numbers
	f, err := os.ReadFile(filename)
	freshCount := 0
	var output = make(map[int]int)
	// var ranges []string
	// var numbers []int
	if err != nil {
		log.Fatal(err)
	}

	file := string(f)

	fileArray := strings.SplitSeq(file, "\n")

	for current := range fileArray {
		if strings.Contains(current, "-") {
			// if this is a range of numbers, take the string and split it.
			twoRanges := strings.Split(current, "-")
			rangeOne, err := strconv.Atoi(twoRanges[0])
			if err != nil {
				log.Fatal(err)
			}
			rangeTwo, err := strconv.Atoi(twoRanges[1])
			if err != nil {
				log.Fatal(err)
			}

			for i := range rangeTwo - rangeOne + 1 {
				// These should be numbers
				// fmt.Println("range", rangeOne+i)
				output[rangeOne+i] = 1
			}

		} else if current != "" {
			n, err := strconv.Atoi(current)
			if err != nil {
				log.Fatal(err)
			}

			// numbers = append(numbers, n)
			_, ok := output[n]
			if ok {
				freshCount++
			}

		}
	}

	// sort.Ints(numbers)
	// fmt.Println(len(numbers))

	// for _, r := range ranges {
	// 	// Convert the string into two numbers
	// 	twoRanges := strings.Split(r, "-")
	// 	rangeOne, err := strconv.Atoi(twoRanges[0])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rangeTwo, err := strconv.Atoi(twoRanges[1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for i := range rangeTwo - rangeOne + 1 {
	// 		// These should be numbers
	// 		// fmt.Println("range", rangeOne+i)
	// 		output[rangeOne+i] = 1
	// 	}

	// 	// fmt.Println("range", r)
	// }

	// fmt.Println("Sorted output.")
	// sort.Ints(numbers)

	// for _, n := range numbers {
	// 	_, ok := output[n]
	// 	if ok {
	// 		freshCount++
	// 	}

	// }
	fmt.Println("Fresh count = ", freshCount)
	return freshCount
}
func main() {
	fmt.Printf("Hello World.")
}
