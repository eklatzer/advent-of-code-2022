package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"strings"

	"advent-of-code-2022/7/util"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

const totalDiskSpace = 70000000
const neededDiskSpaceForUpdate = 30000000

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

				currentFolder = fileSystem.GetSubfolderAndCreateIfNotExists("/")
				if currentLocation != "/" {
					currentFolder = currentFolder.Subfolders.GetSubfolderRecursively(strings.Split(currentLocation[1:], "/"))
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			currentFolder.Subfolders.GetSubfolderAndCreateIfNotExists(commandParts[1])
		} else {
			currentFolder.Files[commandParts[1]] = helpers.ParseInt(commandParts[0])
		}
	}

	totalDiskUsed := generateResult("", fileSystem)
	freeDiskSpace := totalDiskSpace - totalDiskUsed
	neededDiskSpace := neededDiskSpaceForUpdate - freeDiskSpace

	fmt.Println("----")
	var sizeOfFolderToDelete = math.MaxInt
	for path, size := range valuePerPath {
		if size >= neededDiskSpace && size < sizeOfFolderToDelete {
			sizeOfFolderToDelete = size
			fmt.Printf("%s could be deleted (size: %d)\n", path, size)
		}
	}
	log.Printf("size of folder to delete: %d\n", sizeOfFolderToDelete)
	log.Printf("sum of total sizes for directories < 100000: %d", testSum)
}

var testSum = 0
var valuePerPath = map[string]int{}

func generateResult(parentPath string, fileSystem util.FileSystem) int {
	var totalSize = 0
	for path, folderContent := range fileSystem {
		fileSize := folderContent.Files.TotalFileSize()
		var currentPath = filepath.Join(parentPath, path)
		subSize := generateResult(currentPath, *folderContent.Subfolders)
		if (fileSize + subSize) <= 100000 {
			testSum += (fileSize + subSize)
		}
		valuePerPath[currentPath] = (fileSize + subSize)
		fmt.Printf("size for %s: %d (files: %d, sub: %d)\n", currentPath, (fileSize + subSize), fileSize, subSize)
		totalSize += fileSize + subSize
	}
	return totalSize
}
