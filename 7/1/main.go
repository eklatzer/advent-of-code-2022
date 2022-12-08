package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

var testSum = 0

func generateResult(pathBefore string, fileSystem fileSystem) int {
	var totalSize = 0
	for path, folderContent := range fileSystem {
		fileSize := totalFileSize(folderContent.files)
		var currentPath = ""
		if pathBefore == "" {
			currentPath = fmt.Sprintf("%s", path)
		} else {
			currentPath = fmt.Sprintf("%s/%s", pathBefore, path)
		}
		subSize := generateResult(currentPath, *folderContent.subfolders)
		//fmt.Printf("size for %s: %d (files: %d, sub: %d)\n", currentPath, (fileSize + subSize), fileSize, subSize)
		if (fileSize + subSize) <= 100000 {
			testSum += (fileSize + subSize)
			fmt.Printf("size for %s: %d (files: %d, sub: %d)\n", currentPath, (fileSize + subSize), fileSize, subSize)
		}
		totalSize += fileSize + subSize
	}
	return totalSize
}

func totalFileSize(files filesInSystem) int {
	var sum = 0
	for _, size := range files {
		sum += size
	}
	return sum
}

type fileSystem map[string]*folderWithFiles

func (f *fileSystem) getSubfolderRecursivelyFromRoot(paths []string) *folderWithFiles {
	if len(paths) == 1 {
		return (*f).getSubfolderAndCreateIfNotExists(paths[0])
	}
	return (*f)[paths[0]].subfolders.getSubfolderRecursivelyFromRoot(paths[1:])
}

func (f *fileSystem) getSubfolderAndCreateIfNotExists(subpath string) *folderWithFiles {
	if (*f)[subpath] == nil {
		(*f)[subpath] = &folderWithFiles{
			files:      filesInSystem{},
			subfolders: &fileSystem{},
		}
	}
	return (*f)[subpath]
}

type folderWithFiles struct {
	files      filesInSystem
	subfolders *fileSystem
}

type filesInSystem map[string]int

func main() {
	scanner, file, err := helpers.GetInput(helpers.GetInputFilePath())
	defer file.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	var fileSystem = fileSystem{}

	var currentLocation = "/"
	var currentSubfolder = fileSystem.getSubfolderAndCreateIfNotExists("/")

	for scanner.Scan() {
		line := scanner.Text()
		var commandParts = strings.Split(line, " ")

		if strings.HasPrefix(line, "$") {
			if commandParts[1] == "cd" {
				newLocation := commandParts[2]
				if newLocation == "/" {
					currentLocation = "/"
					currentSubfolder = fileSystem.getSubfolderAndCreateIfNotExists("/")
				} else if newLocation == ".." {
					pathParts := strings.Split(currentLocation[1:], "/")
					currentLocation = fmt.Sprintf("/%s", strings.Join(pathParts[:len(pathParts)-1], "/"))
					root := fileSystem.getSubfolderAndCreateIfNotExists("/")
					if currentLocation == "/" {
						currentSubfolder = root
					} else {
						currentSubfolder = root.subfolders.getSubfolderRecursivelyFromRoot(strings.Split(currentLocation[1:], "/"))
					}
				} else {
					currentLocation = filepath.Join(currentLocation, newLocation)
					if currentLocation == "/" {
						currentSubfolder = fileSystem.getSubfolderAndCreateIfNotExists("/")
					} else {
						currentSubfolder = currentSubfolder.subfolders.getSubfolderAndCreateIfNotExists(newLocation)
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			// Handle dir
			currentSubfolder.subfolders.getSubfolderAndCreateIfNotExists(commandParts[1])
		} else {
			// handle file
			log.Println(line)
			currentSubfolder.files[commandParts[1]] = parseInt(commandParts[0])
		}
	}

	log.Println(generateResult("", fileSystem))
	log.Println("result:")
	log.Println(testSum)
}

func parseInt(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}
