package main 

import "fmt"
import "example.com/go/reader"
import "example.com/go/file"

func main() {

	filePaths, err := reader.Read("assets/file.txt")

	if err != nil {
		fmt.Println(err)
	}

	sortedFileSize, err := file.GetSortedFileSize(filePaths)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sortedFileSize)

}