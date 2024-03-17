package text

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/golangast/sugargen/loggers"
)

func GetPropDatatype(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}

func GetField(fields string) ([]string, []string) {
	var field []string
	var property []string
	s := strings.Split(fields, " ")
	for i, ss := range s {
		if i%2 == 0 {
			//get even
			field = append(field, TrimDotright(ss))
		} else {
			property = append(property, TrimDotright(ss))

		}

	}
	return field, property
}

func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}

func UpdateText(f string, o string, n string) error {
	input, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = os.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}
func UpdateTexts(file, check, comment, replace string) error {
	if FindTextNReturn(file, check) != comment {
		err := UpdateText(file, comment, replace+"\n"+comment)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveText(file, check, comment, replace string) error {
	if FindTextNReturn(file, check) != comment {
		err := UpdateText(file, comment, replace+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCode(file, check, comment, replace string) error {
	if FindTextNReturn(file, check) != comment {
		err := UpdateText(file, check, replace+"\n"+comment)
		if err != nil {
			return err
		}
	}
	return nil
}

func Checklogger(err error, s string) {
	logger := loggers.CreateLogger()
	if err != nil {
		logger.Error(
			s,
			slog.String("error: ", err.Error()),
		)
	}
}

// f is for file, o is for old text, n is for new text

func FindTextNReturn(p, str string) string {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	toplevel := TrimDot(str)
	property := TrimDotright(str)
	strs := strings.Replace(property, ".", " ", 1)
	// fmt.Println(str)
	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		if strings.Contains(string(text), toplevel) {
			//is the dot string and split it
			if strings.Contains(string(text), strs) {
				return string(text)
			}
		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			fmt.Println(err)

		}
	}

	// fmt.Println("Reading from file.")
	fmt.Println(string(text))

	return ""
}
