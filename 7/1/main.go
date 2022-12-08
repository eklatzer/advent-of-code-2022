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

var testSum = 0

func generateResult(pathBefore string, fileSystem util.FileSystem) int {
	var totalSize = 0
	for path, folderContent := range fileSystem {
		fileSize := folderContent.Files.TotalFileSize()
		var currentPath = ""
		if pathBefore == "" {
			currentPath = fmt.Sprintf("%s", path)
		} else {
			currentPath = fmt.Sprintf("%s/%s", pathBefore, path)
		}
		subSize := generateResult(currentPath, *folderContent.Subfolders)
		if (fileSize + subSize) <= 100000 {
			testSum += (fileSize + subSize)
			fmt.Printf("size for %s: %d (files: %d, sub: %d)\n", currentPath, (fileSize + subSize), fileSize, subSize)
		}
		totalSize += fileSize + subSize
	}
	return totalSize
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
				newLocation := commandParts[2]
				if strings.HasPrefix(newLocation, "/") {
					currentLocation = newLocation
				} else if newLocation == ".." {
					pathParts := strings.Split(currentLocation[1:], "/")
					currentLocation = fmt.Sprintf("/%s", strings.Join(pathParts[:len(pathParts)-1], "/"))
				} else {
					currentLocation = filepath.Join(currentLocation, newLocation)
				}

				currentFolder = fileSystem.GetSubfolderAndCreateIfNotExists("/")
				if currentLocation != "/" {
					currentFolder = currentFolder.Subfolders.GetSubfolderRecursivelyFromRoot(strings.Split(currentLocation[1:], "/"))
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			currentFolder.Subfolders.GetSubfolderAndCreateIfNotExists(commandParts[1])
		} else {
			currentFolder.Files[commandParts[1]] = helpers.ParseInt(commandParts[0])
		}
	}

	log.Println(generateResult("", fileSystem))
	log.Println("result:")
	log.Println(testSum)
}
