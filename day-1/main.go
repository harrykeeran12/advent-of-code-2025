package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() int {
	f, err := os.Open("day-1/input2.txt")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	// The dial starts at 50.
	var dial = 50

	fmt.Println("The dial starts by pointing to 50.")

	r := bufio.NewReader(f)

	count := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		dir := string(line[0])
		length := len(line)
		distance, err := strconv.Atoi(string(line[1 : length-1]))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(dir, distance)

		if dir == "L" {
			dial = dial - distance
		} else {
			dial = dial + distance
		}
		dial = ((dial % 100) + 100) % 100

		if dial > 99 || dial < 0 {
			log.Fatal(fmt.Sprintf("The dial is at the number %d after distance %d and direction %s\n", dial, distance, dir))
		}

		if dial == 0 {
			count++
		}
		fmt.Printf("The dial is rotated %s %d to point at %d\n", dir, distance, dial)
		// fmt.Printf("Current dial is set to %d\n", dial)
	}
	fmt.Printf("Count = %d \n", count)
	defer f.Close()
	return count

}

func main() {
	f, err := os.Open("day-1/input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// The dial starts at 50.
	var dial = 50

	fmt.Println("The dial starts by pointing to 50.")

	r := bufio.NewReader(f)

	count := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		dir := string(line[0])
		length := len(line)
		distance, err := strconv.Atoi(string(line[1 : length-1]))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(dir, distance)
		//
		// If the distance is bigger than the difference between the dial's position and 0, then you will go over 0.
		//
		// If dial at 49, the distance between it and 0 is 49 to the left and 52 to the right
		// diffL := int(math.Abs(float64(dial)))
		// diffR := 100 - dial
		// prevDial := dial
		// What if the distance is really big? like bigger than 100? Meaning it will take multiple rounds?
		if distance > 100 {
			fmt.Println("Bigger than 100 so add the integer division of distance.")
			divDist := distance / 100
			fmt.Println(divDist)
			count += divDist
			distance = distance - (divDist * 100)
		}
		if dir == "L" {
			difference := dial
			if difference < distance && dial != 0 {
				count++
				fmt.Print("Passes 0 on the left.\n")
			}
			dial = dial - distance
		} else {
			difference := 100 - dial
			if difference < distance && dial != 0 {
				count++
				fmt.Print("Passes 0 on the right.\n")
			}
			dial = dial + distance
		}
		// Normalise here
		dial = ((dial % 100) + 100) % 100
		fmt.Printf("The dial is rotated %s %d to point at %d\n", dir, distance, dial)
		// Take the integer division of the distance with 100, and add that to the count for 0.

		if dial > 99 || dial < 0 {
			log.Fatal(fmt.Sprintf("The dial is at the number %d after distance %d and direction %s\n", dial, distance, dir))
		}

		if dial == 0 {
			count++
		}
		// fmt.Printf("Current dial is set to %d\n", dial)
	}
	//Answer is somewhere between 6000 and 7000
	fmt.Printf("Count = %d \n", count)

	defer f.Close()
}
