package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// A number is an invalid id if the length of the number is repeated.
func isInvalid(id int) bool {
	// Check if id has an odd or even length
	strid := strconv.Itoa(id)
	strlen := len(strid)
	if strlen%2 == 0 {
		midpoint := strlen / 2
		a := strid[0:midpoint]
		b := strid[midpoint:strlen]
		// fmt.Println(strid, "has a length of", strlen)
		if a == b {
			fmt.Println(strid, "is invalid.")
			return true
		}
	}
	return false
}

// A number is an invalid id if the length of the number is repeated twice or more.
func isNewInvalid(id int) bool {
	// Concatenate the string together with itself.
	strid := strconv.Itoa(id)
	strlen := len(strid)
	bigstring := strid + strid
	bigstring = bigstring[1 : (strlen*2)-1]
	// fmt.Println(strid, bigstring[1:strlen * 2 - 1])
	ans := strings.Index(bigstring, strid)
	if ans != -1 {
		fmt.Println(strid, ans+1)
		return true
	}
	return false
}

func part_1(filename string) int {
	// fmt.Println(filename)
	var sum int = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	read := bufio.NewReader(file)

	defer file.Close()
	for {
		line, err := read.ReadString(',')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		line = strings.ReplaceAll(line, ",", "")
		splitted := strings.Split(line, "-")
		first, err := strconv.Atoi(splitted[0])
		if err != nil {
			log.Fatal(err)
		}
		last, err := strconv.Atoi(splitted[1])
		last++
		if err != nil {
			log.Fatal(err)
		}
		// Get the difference between the two ranges.
		diff := int(math.Abs(float64(first - last)))
		// fmt.Println(first, "->", last, diff)
		for i := range diff {
			checkNumber := first + i
			if isInvalid(checkNumber) {
				sum += checkNumber
			}

		}

	}
	fmt.Println("Sum =", sum)
	return sum

}

func part_2(filename string) int {
	// fmt.Println(filename)
	var sum int = 0
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	reading := string(file)
	readings := strings.Split(reading, ",")
	// defer file.Close()
	for _, line := range readings {
		// fmt.Println(item)
		// line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "\n", "")
		splitted := strings.Split(line, "-")
		first, err := strconv.Atoi(splitted[0])
		if err != nil {
			log.Fatal(err)
		}
		last, err := strconv.Atoi(splitted[1])
		last++
		if err != nil {
			log.Fatal(err)
		}
		// Get the difference between the two ranges.
		diff := int(math.Abs(float64(first - last)))
		fmt.Println(first, "->", last, diff)
		for i := range diff {
			checkNumber := first + i
			if isNewInvalid(checkNumber) {
				sum += checkNumber
			}
		}
	}

	// read := bufio.NewReader(file)
	// for {
	// 	// line, err := read.ReadString(',')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	line = strings.ReplaceAll(line, ",", "")
	// 	splitted := strings.Split(line, "-")
	// 	first, err := strconv.Atoi(splitted[0])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	last, err := strconv.Atoi(splitted[1])
	// 	last++
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// Get the difference between the two ranges.
	// 	diff := int(math.Abs(float64(first - last)))
	// 	fmt.Println(first, "->", last, diff)
	// 	for i := range diff {
	// 		checkNumber := first + i
	// 		if isNewInvalid(checkNumber) {
	// 			sum += checkNumber
	// 		}

	// 	}

	// }
	fmt.Println("Sum =", sum)
	return sum

}

func main() {
	// fmt.Printf("Hello World.")
	// part_1("day-2/input.txt")
	// part_1("day-2/input2.txt")
	part_2("day-2/input.txt")

}
