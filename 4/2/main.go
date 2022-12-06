package main

import (
	"flag"

	"advent-of-code-2022/4/util"
)

func init() {
	flag.Parse()
}

func main() {
	util.Execute(func(sectionElveOne, sectionElveTwo util.SectionRange) bool {
		return sectionElveOne.OverlapsWith(sectionElveTwo)
	})
}
