package main

import (
	"fmt"
	"slices"

	"github.com/golangast/sugargen/bash"
	"github.com/golangast/sugargen/chat"
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
	text.UpdateTexts(answer+"/"+coloranswer+".txt", "template map", "template map", "template cat")

	//now lets spit out whats in the file using bach in go.
	if err := bash.ShellBash("cat ./" + answer + "/" + coloranswer + ".txt"); err != nil {
		text.Checklogger(err, "trying to do bash in go")
	}

	//start training the model! specify the model name and the filename
	chat.CheckIfSpanLimitsEqualText("server", "server.json")

	//use an input to ask a question
	ans := input.InputScanDirections("What would you like to do?")
	//get data from the model
	text, label := chat.GetTextLabelFromGlob(ans)
	fmt.Println(text)
	fmt.Println(label)

	//start to use that data to run commands!
	if slices.Contains(label, "server") {
		fmt.Println("Starting the server...")
	}
	//a training file is required and it does have to have pre-processed data
	//in a particular format.
	/*
		[
			{
			    "Text": "sentence that contains the text or phrase",
			    "Spans": [
			      {
			        "Start": 13, //where the word or phrase starts
			        "End": 28, //where the word or phrase ends
			        "Label": "server" //name of the mobel
			      }
			    ],
			    "Answer": "reject" //whether it will be accepted or rejected
			  },
			]
			  technically you can add spans and create more models and
			  change the format but you will need to update the data structure

			  right now we are going to use /server.json
			  The name of the model does matter
	*/
	//now you have the ability to get user input two different ways and generate files, update text, use bash,
	//and train a model and use that data for commands

	//the rest is up to you!
}
