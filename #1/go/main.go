package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var calories = []int{}
	var currentCalories = 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		c, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("failed to convert line to int: %v\n", err)
			continue
		}
		currentCalories += c
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	var totalCaloriesForTopElves = 0
	for i := 0; i < 3; i++ {
		log.Printf("#%d: %d calories", i+1, calories[i])
		totalCaloriesForTopElves += calories[i]
	}
	log.Print("----")
	log.Printf("total calories of top elves: %d calories", totalCaloriesForTopElves)
}
