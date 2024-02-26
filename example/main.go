package main

import (
	"fmt"

	"github.com/golangast/sugargen/bash"
	"github.com/golangast/sugargen/filefolder"
	"github.com/golangast/sugargen/gen"
	"github.com/golangast/sugargen/input"
	"github.com/golangast/sugargen/text"
)

func main() {
	// Parse command line input for a question
	answer := input.InputScanDirections("what is your name?")
	fmt.Println(answer, " nice name")

	// show a list of selection and chosen one is stored
	var colorsslice = []string{"red", "green", "blue", "yellow"}

	//remember you can set the number from 1 to whatever you want.  1 is the number of columns
	//to add colors update the colors/colors.go file and use https://en.wikipedia.org/wiki/ANSI_escape_code#colors
	coloranswer := input.MenuInstuctions(colorsslice, 1, "purple", "purple", "which color do you prefer to use?")
	fmt.Println(coloranswer, " is the right answer")

	//what can you do with all this? Plenty!
	//how about lets make a folder file from the answers
	if err := filefolder.Makefolder(answer); err != nil {
		text.Checklogger(err, "trying to create folder")
	}

	file, err := filefolder.Makefile(answer + "/" + coloranswer + ".txt")
	text.Checklogger(err, "trying to add text to file")

	//you wanna update text in it? Sure we can do that.
	//we need some text first.
	var sometext = `This is some text but lets add a template map {{.name}}`

	//create a map
	mb := make(map[string]string)
	mb["name"] = answer

	//now lets do some generating
	if err := gen.Writetemplate(sometext, file, mb); err != nil {
		text.Checklogger(err, "trying to create folder")
	}

	//now lets try to update the text "template map" in the file to "template cat".
	text.UpdateText(answer+"/"+coloranswer+".txt", "template map", "template map", "template cat")

	//now lets spit out whats in the file using bach in go.
	if err := bash.ShellBash("cat ./" + answer + "/" + coloranswer + ".txt"); err != nil {
		text.Checklogger(err, "trying to do bash in go")
	}
	//now you have the ability to get user input two different ways and generate files, update text, and use bash

	//the rest is up to you!
}
