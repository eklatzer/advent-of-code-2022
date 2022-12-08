package util

type FileSystem map[string]*DirectoryWithFiles

type DirectoryWithFiles struct {
	Files      listOfFiles
	Subfolders *FileSystem
}

type listOfFiles map[string]int

func (l listOfFiles) TotalFileSize() int {
	var sum = 0
	for _, size := range l {
		sum += size
	}
	return sum
}

func (f *FileSystem) GetSubfolderRecursivelyFromRoot(paths []string) *DirectoryWithFiles {
	if len(paths) == 1 {
		return (*f).GetSubfolderAndCreateIfNotExists(paths[0])
	}
	return (*f)[paths[0]].Subfolders.GetSubfolderRecursivelyFromRoot(paths[1:])
}

func (f *FileSystem) GetSubfolderAndCreateIfNotExists(subpath string) *DirectoryWithFiles {
	if (*f)[subpath] == nil {
		(*f)[subpath] = &DirectoryWithFiles{
			Files:      make(map[string]int),
			Subfolders: &FileSystem{},
		}
	}
	return (*f)[subpath]
}
