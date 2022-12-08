package util

import (
	"path/filepath"
	"strings"
)

type FileSystem map[string]*DirectoryWithFiles

type DirectoryWithFiles struct {
	Files      listOfFiles
	Subfolders *FileSystem
}

func newDirectoryWithFiles() *DirectoryWithFiles {
	return &DirectoryWithFiles{
		Files:      make(map[string]int),
		Subfolders: &FileSystem{},
	}
}

type listOfFiles map[string]int

func (l listOfFiles) TotalFileSize() int {
	var sum = 0
	for _, size := range l {
		sum += size
	}
	return sum
}

func (f *FileSystem) GetSubfolderRecursively(paths []string) *DirectoryWithFiles {
	if len(paths) == 1 {
		return (*f).GetSubfolderAndCreateIfNotExists(paths[0])
	}
	return (*f)[paths[0]].Subfolders.GetSubfolderRecursively(paths[1:])
}

func (f *FileSystem) GetSubfolderAndCreateIfNotExists(subpath string) *DirectoryWithFiles {
	if (*f)[subpath] == nil {
		(*f)[subpath] = newDirectoryWithFiles()
	}
	return (*f)[subpath]
}

func GetNewTotalPath(currentLocation, newLocation string) string {
	switch {
	case strings.HasPrefix(newLocation, "/"):
		return newLocation
	case newLocation == "..":
		indexLastSlash := strings.LastIndex(currentLocation, "/")
		return filepath.Join("/", currentLocation[:indexLastSlash])
	}
	return filepath.Join(currentLocation, newLocation)
}
