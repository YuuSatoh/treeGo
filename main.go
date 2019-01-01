package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var hasPaentOtherFiles []bool

func main() {
	dirName, fileName := initInput()
	runTree(dirName, fileName)
}

func runTree(dir string, file string) {
	readDirPath := dir + file
	readedfiles, err := ioutil.ReadDir(readDirPath)
	if err != nil {
		fmt.Printf("%s: %s", os.Args[0], err)
		os.Exit(1)
	}

	readedFilesLength := len(readedfiles)
	for index, rededFile := range readedfiles {

		if strings.HasPrefix(rededFile.Name(), ".") {
			continue
		}

		isLastFile := isLast(index, readedFilesLength)
		printTreeLine(rededFile.Name(), isLastFile)

		if rededFile.IsDir() {
			hasPaentOtherFiles = append(hasPaentOtherFiles, isLastFile)
			runTree(readDirPath+"/", rededFile.Name())
		}
	}
	hasPaentOtherFiles = deleteSliceLastItem(hasPaentOtherFiles)
}

func isLast(index int, len int) bool {
	return index == len-1
}

func printTreeLine(fileName string, isLast bool) {
	for _, hasPaentOtherFile := range hasPaentOtherFiles {
		if hasPaentOtherFile {
			fmt.Print("    ")
		} else {
			fmt.Print("│   ")
		}
	}

	if isLast {
		fmt.Printf("└─")
	} else {
		fmt.Printf("├─")
	}

	fmt.Println("─", fileName)
}

func deleteSliceLastItem(slice []bool) []bool {
	if slice == nil || len(slice) < 1 {
		return slice
	}
	tempSlice := slice[:len(slice)-1]
	newSlice := make([]bool, len(tempSlice))
	copy(newSlice, slice)
	return newSlice
}

func initInput() (string, string) {
	var (
		path string
		dir  string
		file string
	)
	flag.StringVar(&path, "path", "", "print dir path flag")
	flag.Parse()
	dir, file = filepath.Split(path)
	if dir == "" {
		dir, _ = os.Getwd()
	}
	fmt.Println(dir)
	return dir, file
}
