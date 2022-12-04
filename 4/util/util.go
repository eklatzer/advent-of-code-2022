package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent-of-code-2022/helpers"
)

// Execute reads the file, extracts the ranges and counts the ranges meeting the condition.
func Execute(condition func(SectionRange, SectionRange) bool) {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var numberOfPairsForCondition = 0

	for scanner.Scan() {
		sectionElveOne, sectionElveTwo, err := extractRanges(scanner.Text())
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if condition(sectionElveOne, sectionElveTwo) {
			numberOfPairsForCondition++
		}
	}
	log.Printf("number of pairs meeting the requirement: %d", numberOfPairsForCondition)
}

// SectionRange defines a range.
type SectionRange struct {
	start int
	end   int
}

// FromString sets the values of a SectionRange from a string in the format 'min-max' (e.g. '10-50')
func (s *SectionRange) FromString(in string) error {
	sectionParts := strings.Split(in, "-")
	if len(sectionParts) != 2 {
		return fmt.Errorf("invalid input, expected two numbers separated by hyphen, got: %s", in)
	}
	var err error

	s.start, err = strconv.Atoi(sectionParts[0])
	if err != nil {
		return fmt.Errorf("failed to cast start of range %q to int", sectionParts[0])
	}
	s.end, err = strconv.Atoi(sectionParts[1])
	if err != nil {
		return fmt.Errorf("failed to cast end of range %q to int", sectionParts[1])
	}
	return nil
}

// FullyContains checks if s fully contains s2.
func (s *SectionRange) FullyContains(s2 SectionRange) bool {
	return s.start <= s2.start && s.end >= s2.end
}

// OverlapsWith checks if s and s2 overlap.
func (s *SectionRange) OverlapsWith(s2 SectionRange) bool {
	return s.start <= s2.end && s.end >= s2.start
}

func extractRanges(line string) (sectionElveOne SectionRange, sectionElveTwo SectionRange, err error) {
	rangeDefinitions := strings.Split(line, ",")
	if len(rangeDefinitions) != 2 {
		return sectionElveOne, sectionElveTwo, fmt.Errorf("invalid line, expected two ranges separated by comma, got: %s", line)
	}

	err = sectionElveOne.FromString(rangeDefinitions[0])
	if err != nil {
		return sectionElveOne, sectionElveTwo, err
	}

	err = sectionElveTwo.FromString(rangeDefinitions[1])
	if err != nil {
		return sectionElveOne, sectionElveTwo, err
	}
	return sectionElveOne, sectionElveTwo, nil
}
