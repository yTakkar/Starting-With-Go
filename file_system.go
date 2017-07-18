package main

import (
	"log"
	"os"
)

func main() {

	// Create file
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Write to a file
	file.WriteString("Hello, world!!")
	file.Close()

	// Check if a file exists, if not make a directory
	if _, stat_err := os.Stat("./folder"); os.IsNotExist(stat_err) {
		mkerr := os.Mkdir("folder", 0655)
		if mkerr != nil {
			log.Fatal(mkerr)
		}
	}

	// Copy file
	copy_err := os.Link("./file.txt", "./folder/newfile.txt")
	if copy_err != nil {
		log.Fatal(copy_err)
	}

	// Rename
	rename_err := os.Rename("./file.txt", "./renamed_file.txt")
	if rename_err != nil {
		log.Fatal(rename_err)
	}

	// Delete a file or a folder
	rem_err := os.Remove("./renamed_file.txt")
	if rem_err != nil {
		log.Fatal(err)
	}

	// Delete a folder and all its contents (if exists)
	rema_err := os.RemoveAll("./folder")
	if rema_err != nil {
		log.Fatal(rema_err)
	}

}
