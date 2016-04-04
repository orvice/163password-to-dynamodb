package saver

import (
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"log"
)

func genFilePath(path,file string) string {
	return path + file;
}

func processFiles(path string){
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		file, err := os.Open(genFilePath(path,f))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}