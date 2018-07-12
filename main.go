package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {

	flags := parseFlags()

	home := os.Getenv("HOME")
	noteFolder := home + "/.gonote/"

	filesExist := checkFiles(noteFolder)
	if !filesExist {

		if err := createFiles(noteFolder); err != nil {
			fmt.Println(color.RedString("ERROR: ") + "Unable to create files in " + noteFolder)
			panic(err)
		}
	}

	if *flags.shouldView {
		viewNotes(noteFolder)
	} else if *flags.delete != (-1) {
		deleteNote(*flags.delete, noteFolder)
	} else if *flags.content != "" {
		newNote(*flags.content).add(noteFolder)
	} else {
		viewNotes(noteFolder)
	}
}
