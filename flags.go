package main

import (
	"flag"
)

type flagResult struct {
	content    *string
	shouldView *bool
	delete     *int
}

func parseFlags() *flagResult {

	res := new(flagResult)

	res.content = flag.String("n", "", "Create new note")
	res.shouldView = flag.Bool("v", false, "View existing notes")
	res.delete = flag.Int("d", (-1), "Delete an note")
	flag.Parse()

	return res
}
