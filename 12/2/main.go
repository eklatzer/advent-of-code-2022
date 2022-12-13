package main

import (
	"flag"
	"log"
	"sort"

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

	var starts []*graph.Node[util.Position, byte]
	var destination *graph.Node[util.Position, byte]
	for _, n := range nodes {
		if n.Value == 'S' || n.Value == 'a' {
			n.Value = 'a'
			starts = append(starts, n)
		} else if n.Value == 'E' {
			n.Value = 'z'
			destination = n
		}
	}

	var paths = map[util.Position]int{}
	log.Println(len(starts))

	for _, start := range starts {
		steps := len(util.FindShortestPath(nodes, start, destination))
		if steps != 0 {
			paths[start.Identifier] = steps
		}
	}

	keys := make([]util.Position, 0, len(paths))

	for k := range paths {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return paths[keys[i]] < paths[keys[j]]
	})

	for _, position := range keys {
		log.Printf("%v: %d", position, paths[position])
	}

}
