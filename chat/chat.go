package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/jdkato/prose/v2"
)

// seperator for good and bad data
func CollectAcceptedRejected(data []ProdigyOutput) ([]prose.EntityContext, []ProdigyOutput) {
	cutoff := int(float64(len(data)) * 0.8) // cutoff for looking through data

	//initialize
	train, test := []prose.EntityContext{}, []ProdigyOutput{}

	//separate data where one is accepted
	for i, entry := range data {
		if i < cutoff {
			//simplify data to struct
			train = append(train, prose.EntityContext{Text: entry.Text, Spans: entry.Spans, Accept: entry.Answer == "accept"})
		} else {
			test = append(test, entry)
		}
	}

	return train, test
}

// get data and turn it into a model blob
func CheckIfSpanLimitsEqualText(models, file string) {
	//clear model blob
	err := os.RemoveAll("Maxent")
	if err != nil {
		fmt.Println(err)
	}
	//get data
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("please make a json file that has training data.  One can be found here https://github.com/golangast/sugargen/blob/main/example/server.json")
		panic(err)
	}
	//turn into a model blob
	decodedata, err := DecodeJSONToProdigyOutput(data)
	if err != nil {
		panic(err)
	}
	//test is rejected and train is accepted
	train, test := CollectAcceptedRejected(decodedata)

	//create model
	model := prose.ModelFromData(models, prose.UsingEntities(train))

	//turn model into a doc type
	correct := 0.0
	for _, entry := range test {
		doc, err := prose.NewDocument(entry.Text, prose.WithSegmentation(false), prose.UsingModel(model))
		if err != nil {
			panic(err)
		}
		ents := doc.Entities()

		if entry.Answer != "accept" && len(ents) == 0 { //bad data
			correct++

		} else { //making sure its good data

			expected := []string{}

			//look for text in span
			for _, span := range entry.Spans {
				expected = append(expected, entry.Text[span.Start:span.End])
			}

			//see if text matches entities
			if reflect.DeepEqual(expected, ents) {
				correct++
			}
		}
	}
	//show results
	fmt.Printf("Correct (%%): %f\n", correct/float64(len(test)))
	//create blob model file
	model.Write(".")
}

// get data from blob model and create entities
func GetTextLabelFromGlob(str string) ([]string, []string) {

	models := prose.ModelFromDisk(".") //get blob model

	//turn it into a doc
	doc, err := prose.NewDocument(str, prose.WithSegmentation(false), prose.UsingModel(models))

	if err != nil {
		panic(err)
	}
	var entstext []string
	var entslabel []string
	// get entities labels
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		entstext = append(entstext, ent.Text)
		entslabel = append(entslabel, ent.Label)

	}

	return entstext, entslabel

}

// get data from json file for training
func DecodeJSONToProdigyOutput(jsonData []byte) ([]ProdigyOutput, error) {
	entries := []ProdigyOutput{}

	err := json.Unmarshal(jsonData, &entries)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, err
	}

	return entries, nil

}

type ProdigyOutput struct {
	Text   string
	Spans  []prose.LabeledEntity
	Answer string
}
