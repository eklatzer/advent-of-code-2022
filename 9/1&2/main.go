package main

import (
	"flag"
	"log"

	"advent-of-code-2022/9/util"
	"advent-of-code-2022/helpers"
)

var numberOfKnots int

func init() {
	flag.IntVar(&numberOfKnots, "n", 2, "Number of knots") // part 1: n=2, part 2: n=10
	flag.Parse()
}

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	positions := make([]util.Position, numberOfKnots)

	field := helpers.Set[util.Position]{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		cmd := util.ExtractCommand(line)

		for ; cmd.Distance > 0; cmd.Distance-- {
			switch cmd.Direction {
			case util.Up:
				positions[0].Y++
			case util.Down:
				positions[0].Y--
			case util.Right:
				positions[0].X++
			case util.Left:
				positions[0].X--
			}
			for i := 1; i < len(positions); i++ {
				positions[i].Follow(positions[i-1])
			}
			field[positions[len(positions)-1]] = struct{}{}
		}
	}

	log.Printf("%d positions are visited by the tail at least once", len(field))
}
