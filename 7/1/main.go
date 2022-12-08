package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"advent-of-code-2022/7/util"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var fileSystem = util.FileSystem{}

	var currentLocation = "/"
	var currentFolder = fileSystem.GetSubfolderAndCreateIfNotExists("/")

	for scanner.Scan() {
		line := scanner.Text()
		var commandParts = strings.Split(line, " ")

		if strings.HasPrefix(line, "$") {
			if commandParts[1] == "cd" {
				currentLocation = util.GetNewTotalPath(currentLocation, commandParts[2])
				currentFolder = fileSystem.GetSubfolderAndCreateIfNotExists("/").Subfolders.GetSubfolderRecursively(strings.Split(currentLocation[1:], "/"))
			}
		} else if strings.HasPrefix(line, "dir") {
			currentFolder.Subfolders.GetSubfolderAndCreateIfNotExists(commandParts[1])
		} else {
			currentFolder.Files[commandParts[1]] = helpers.ParseInt(commandParts[0])
		}
	}

	generateResult("", fileSystem)
	log.Printf("sum of total sizes for directories < 100000: %d", testSum)
}

var testSum = 0

func generateResult(parentPath string, fileSystem util.FileSystem) int {
	var totalSize = 0
	for path, folderContent := range fileSystem {
		fileSize := folderContent.Files.TotalFileSize()
		var currentPath = filepath.Join(parentPath, path)
		subSize := generateResult(currentPath, *folderContent.Subfolders)
		if (fileSize + subSize) <= 100000 {
			testSum += (fileSize + subSize)
			fmt.Printf("size for %s: %d (files: %d, sub: %d)\n", currentPath, (fileSize + subSize), fileSize, subSize)
		}
		totalSize += fileSize + subSize
	}
	return totalSize
}
