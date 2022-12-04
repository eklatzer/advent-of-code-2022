package main

import (
	"advent-of-code-2022/4/util"
)

func main() {
	util.Execute(func(sectionElveOne, sectionElveTwo util.SectionRange) bool {
		return sectionElveOne.OverlapsWith(sectionElveTwo)
	})
}
