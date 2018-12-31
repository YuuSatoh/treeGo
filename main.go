package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"cmd"
)

func main() {

	// 取得
	p, _ := os.Getwd()
	fmt.Println(p)

	// var path string
	path := initOption()
	fmt.Println(path)

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		fmt.Println(f.Name())
	}

    if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

}

func initOption() string {
	var dirPath string
	flag.StringVar(&dirPath, "path", "./", "print dir path flag")
	flag.Parse()
	return dirPath
}


func run() {
    cmd := cmd.NewFoobarCommand()
    return cmd.Execute()
}