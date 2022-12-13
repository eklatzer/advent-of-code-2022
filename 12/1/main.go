package main

import (
	"flag"
	"log"

	"advent-of-code-2022/12/util"
	"advent-of-code-2022/12/util/graph"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	nodes, err := util.GetNodesFromFile(helpers.GetInputFilePath())

	if err != nil {
		log.Fatalln(err)
	}

	var root *graph.Node[util.Position, byte]
	var destination *graph.Node[util.Position, byte]
	for _, n := range nodes {
		if n.Value == 'S' {
			root = n
		} else if n.Value == 'E' {
			n.Value = 'z'
			destination = n
		}
	}

	stepsToDestination := util.FindShortestPath(nodes, root, destination)

	log.Printf("number of steps from start to end: %d", len(stepsToDestination))
}
