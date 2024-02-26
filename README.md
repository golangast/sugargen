## General info
This is a utility for terminals, inputs and generating code.


![GitHub repo file count](https://img.shields.io/github/directory-file-count/golangast/sugargen) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/golangast/sugargen)
![GitHub repo size](https://img.shields.io/github/repo-size/golangast/sugargen)
![GitHub](https://img.shields.io/github/license/golangast/sugargen)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/golangast/sugargen)
![Go 100%](https://img.shields.io/badge/Go-100%25-blue)
![status beta](https://img.shields.io/badge/Status-Beta-red)
<img src="https://img.shields.io/github/license/golangast/sugargen.svg"><img src="https://img.shields.io/github/stars/golangast/sugargen.svg">[![Website shields.io](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](http://endrulats.com)[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/golangast/sugargen)[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/golangast/sugargen)[![GoReportCard example](https://goreportcard.com/badge/github.com/golangast/sugargen)](https://goreportcard.com/report/github.com/golangast/sugargen)[![saythanks](https://img.shields.io/badge/say-thanks-ff69b4.svg)](https://saythanks.io/to/golangast)


## Requirements
* go 1.21.5

## Why build this?
* Go never changes
* Created sugar syntax to allow others and myself to create projects like [Switchterm](https://github.com/golangast/switchterm)


* the following is how you use it and there is an example file in the example folder.
```go
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

```

## Things to remember
* using atomicgo.dev/keyboard there is no way to call itself after a key press

## Licenses
1. [GNU 3 for my code](https://github.com/golangast/switchterm/blob/main/LICENSE.md)
2. [MIT License for atomicgo keyboard](https://github.com/atomicgo/keyboard/blob/main/LICENSE)
2. [MIT License for sprig](https://github.com/Masterminds/sprig?tab=MIT-1-ov-file#readme)
3. [BSD-3-Clause for sqlite driver](https://pkg.go.dev/modernc.org/sqlite?tab=licenses) 
4. [BSD-3-Clause for Go itself](https://github.com/golang/go/blob/master/LICENSE) 



## Special thanks
* [Go Team because they are gods](https://github.com/golang/go/graphs/contributors)
* [Creators of https://pkg.go.dev/modernc.org/sqlite - ](https://gitlab.com/cznic/sqlite/-/project_members)
* [Creators of https://github.com/Masterminds/sprig- ](https://github.com/Masterminds/sprig/graphs/contributors)
* [Creators of https://github.com/atomicgo/keyboard - ](https://github.com/MarvinJWendt)

## Why Go?
* The language is done since 1.0.https://youtu.be/rFejpH_tAHM there are little features that get added after 10 years but whatever you learn now will forever be useful.
* It also has a compatibility promise https://go.dev/doc/go1compat
* It was also built by great people. https://hackernoon.com/why-go-ef8850dc5f3c
* 14th used language https://insights.stackoverflow.com/survey/2021
* Highest starred language https://github.com/golang/go
* It is also number 1 language to go to and not from https://www.jetbrains.com/lp/devecosystem-2021/#Do-you-plan-to-adopt--migrate-to-other-languages-in-the-next--months-If-so-to-which-ones
* Go is growing in all measures https://madnight.github.io/githut/#/stars/2023/3
* Jobs are almost doubling every year. https://stacktrends.dev/technologies/programming-languages/golang/
* Companies that use go. https://go.dev/wiki/GoUsers
* Why I picked Go https://youtu.be/fD005g07cU4
