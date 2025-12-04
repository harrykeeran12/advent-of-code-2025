package day3

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func maxBattery(batteryList []string) (int, int) {
	var maxNum, maxIndex int = -1, -1

	for index, battery := range batteryList {
		// Get the battery, and find the biggest one.
		batteryNum, err := strconv.Atoi(battery)
		if err != nil {
			log.Fatal(err)
		}
		if batteryNum > maxNum {
			// Set maxNum to batterynum
			maxIndex = index
			maxNum = batteryNum
		}
	}

	return maxNum, maxIndex
}

func maxBatteryIgnoreIndex(batteryList []string, ignoreIndex []int) (int, int) {
	var maxNum, maxIndex int = -1, -1

	for index, battery := range batteryList {
		// Get the battery, and find the biggest one.
		batteryNum, err := strconv.Atoi(battery)
		if err != nil {
			log.Fatal(err)
		}
		if !slices.Contains(ignoreIndex, index) {
			if batteryNum > maxNum {
				// Set maxNum to batterynum
				maxIndex = index
				maxNum = batteryNum

			}
		}

	}
	fmt.Println("Index at ", maxIndex, "number is ", maxNum)

	return maxNum, maxIndex
}

func largestJolts(bankString string) int {
	// Get the largest jolts by taking the numbers inside the bank strings and comparing the indexes.
	// Convert each string into a list of numbers.
	fmt.Println(bankString)
	batteryList := strings.Split(bankString, "")
	batteryCount := len(batteryList)
	var maxNum, maxIndex int = -1, -1
	var secondNum, secondIndex int = -1, -1

	maxNum, maxIndex = maxBattery(batteryList)

	// Check if the index is equal to the last digit or not
	if maxIndex == len(batteryList)-1 {
		secondNum, secondIndex = maxBattery(batteryList[0:maxIndex])

	} else {
		// If the max index has numbers after it, take a slice of the max
		secondNum, secondIndex = maxBattery(batteryList[maxIndex+1 : batteryCount])
		// As second index starts from max index so needs to be offset.
		secondIndex += maxIndex
	}
	fmt.Printf("Biggest numbers are %d(%d) and %d(%d)\n", maxNum, maxIndex, secondNum, secondIndex)
	if maxIndex > secondIndex {
		res, err := strconv.Atoi(strconv.Itoa(secondNum) + strconv.Itoa(maxNum))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Answer is ", res)
		return res
	} else {
		res, err := strconv.Atoi(strconv.Itoa(maxNum) + strconv.Itoa(secondNum))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Answer is ", res)
		return res
	}
}

func part1(filename string) int {
	f, err := os.ReadFile(filename)
	maxJoltage := 0
	if err != nil {
		log.Fatal(err)
	}

	file := string(f)

	fileArray := strings.SplitSeq(file, "\n")

	for bank := range fileArray {
		if bank != "" {
			// fmt.Println(bank)
			maxJoltage += largestJolts(bank)
		}

	}

	return maxJoltage
}

func largestTwelveJolts(bankString string) int {
	// Get the largest 12 jolts by taking the numbers inside the bank strings and comparing the indexes.
	// Convert each string into a list of numbers.
	fmt.Println(bankString)
	batteryList := strings.Split(bankString, "")
	// batteryCount := len(batteryList)
	var ignoredIndexes []int
	// Iterate over the numbers
	// Make a hashmap for each index and number.
	hmap := make(map[int]int)
	for range 12 {
		maxVal, ignore := maxBatteryIgnoreIndex(batteryList, ignoredIndexes)
		ignoredIndexes = append(ignoredIndexes, ignore)
		hmap[ignore] = maxVal
	}
	fmt.Println(ignoredIndexes)
	sort.Ints(ignoredIndexes)
	fmt.Println(ignoredIndexes)

	var joltString string
	for _, ignored := range ignoredIndexes{
		fmt.Println(ignored, string(bankString[ignored]))
		joltString += string(bankString[ignored])
	}
	fmt.Println(joltString)

	res, err := strconv.Atoi(joltString)
	if err != nil{
		log.Fatal(err)
}


	return res
}

func part2(filename string) int {
	// Take each element inside the filename, and read the max joltage by
	f, err := os.ReadFile(filename)
	maxJoltage := 0
	if err != nil {
		log.Fatal(err)
	}

	file := string(f)

	fileArray := strings.SplitSeq(file, "\n")

	for bank := range fileArray {
		if bank != "" {
			// fmt.Println(bank)
			maxJoltage += largestTwelveJolts(bank)
		}

	}

	return maxJoltage
}

func main() {
	part1("day-3/input2.txt")
}
