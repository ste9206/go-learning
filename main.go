package main

import "fmt"
import "example.com/go/reader"
import "example.com/go/file"

func main() {

	filePaths, err := reader.Read("assets/file.txt")

	if err != nil {
		fmt.Println(err)
	}

	sortedFileSize := file.GetSortedFileSize(filePaths)

	fmt.Println(sortedFileSize)

}
