package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var dir string
	var processType string
	reader := bufio.NewReader(os.Stdin)

	autoResponse := []string{"auto", "a", "A", "AUTO"}
	manualResponse := []string{"manual", "m", "MANUAL", "M"}
	yesResponses := []string{"y", "Y", "yes", "YES", "Yes", "DELETE"}
	noResponses := []string{"n", "N", "no", "NO", "No"}

	// Get the process type
	fmt.Println("Manual or Auto? (m/a) ")
	processType, _ = reader.ReadString('\n')

	// Get the directory to scan
	fmt.Println("Please type in directory to clean: ")
	dir, _ = reader.ReadString('\n')
	fmt.Printf("Scanning Dir: %s", dir)

	// Read directory for file list
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through files and perform action depending on processType
	for _, f := range files {
		fmt.Println(f.Name())
		if containsString(manualResponse, processType) {
			fmt.Println("Delete File:", f.Name())
			resp, _ := reader.ReadString('\n')
			if containsString(yesResponses, resp) {
				fmt.Println("Deleting File...")
			} else if containsString(noResponses, resp) {
				fmt.Println("Skipping File...")
			}
		} else if containsString(autoResponse, processType) {
			fmt.Printf("Deleting File: %s", f.Name())
		}
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
