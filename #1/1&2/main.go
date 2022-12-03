package main

import (
	"log"
	"sort"
	"strconv"

	"advent-of-code-2022/helpers"
)

func main() {
	scanner, file, err := helpers.GetInput("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

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
