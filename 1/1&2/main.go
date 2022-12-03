package main

import (
	"flag"
	"log"
	"sort"
	"strconv"

	"advent-of-code-2022/helpers"
)

func main() {
	numberOfElves := *flag.Int("n", 3, "Number of top elves to sum up")
	flag.Parse()

	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var caloriesPerElve = make([]int, 1)

	for scanner.Scan() {
		line := scanner.Text()
		c, err := strconv.Atoi(line)
		if err != nil {
			caloriesPerElve = append(caloriesPerElve, 0)
			continue
		}
		caloriesPerElve[len(caloriesPerElve)-1] += c
	}
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElve)))

	var totalCaloriesForTopElves = 0
	for i := 0; i < numberOfElves; i++ {
		log.Printf("#%d: %d calories", i+1, caloriesPerElve[i])
		totalCaloriesForTopElves += caloriesPerElve[i]
	}
	log.Print("------------------------")
	log.Printf("total calories of top %d elves: %d calories", numberOfElves, totalCaloriesForTopElves)
}
