package main

import (
	"encoding/gob"
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

// Note is what gets saved and read
type Note struct {
	// ID is now index instead, much more flexible
	Content string
	Time    string
}

func checkFiles(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFiles(f string) error {

	err := os.Mkdir(f, 0722)
	gobfile, err := os.Create(f + "notes.gob")
	if err != nil {
		return err
	}
	defer gobfile.Close()

	var EmptyNoteList []Note

	encoder := gob.NewEncoder(gobfile)
	err = encoder.Encode(EmptyNoteList)

	return err
}

func openNotes(f string) []Note {

	var Notes []Note

	gobfile, err := os.OpenFile(f+"notes.gob", os.O_RDWR, 0722)
	defer gobfile.Close()
	if err == nil {

		decoder := gob.NewDecoder(gobfile)
		err := decoder.Decode(&Notes)

		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	return Notes
}

func (n *Note) display(i int) {
	fmt.Println(color.GreenString(" %d [", i), color.RedString(n.Time), color.GreenString("]: "), n.Content)
}

func viewNotes(f string) {

	notes := openNotes(f)

	for i, v := range notes {
		v.display(i)
	}

}

func deleteNoteUI() {

}

func newNote(Content string) Note {
	var n Note

	// Get and format current date
	UTCTime := time.Now().UTC()
	n.Time = UTCTime.Format("2006-01-02")

	n.Content = Content

	return n
}

func (n Note) add(f string) {

	notes := openNotes(f)

	notes = append(notes, n)

	encodeNotes(notes, f)

}

func encodeNotes(notes []Note, f string) {

	gobfile, err := os.OpenFile(f+"notes.gob", os.O_RDWR, 0722)
	if err != nil {
		panic(err)
	}
	defer gobfile.Close()
	encoder := gob.NewEncoder(gobfile)

	err = encoder.Encode(notes)
	if err != nil {
		panic(err)
	}

}

func deleteNote(ID int, f string) {

	notes := openNotes(f)

	for i := range notes {
		if i == ID {
			notes = append(notes[:i], notes[i+1:]...)
			encodeNotes(notes, f)
			return
		}
	}
}
