package jtl

import (
	"encoding/xml"
	"fmt"
	"os"
)

func Decode(filename string, output chan<- interface{}, mergedJtlFiles []string) {
	defer close(output)

	WriteToOutputChannel(filename, output, false)
	for i := 0; i < len(mergedJtlFiles); i++ {
		WriteToOutputChannel(mergedJtlFiles[i], output, true)
	}
}

func WriteToOutputChannel(filename string, output chan<- interface{}, needMerge bool) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		// Inspect the type of the token just read.
		switch elementType := token.(type) {
		case xml.StartElement:
			name := elementType.Name.Local
			if name == "testResults" {
				version := elementType.Attr[0].Value
				// send testResult tag
				testResults := TestResults{Version: version, XMLName: xml.Name{Space: "", Local: "testResults"}}
				if !needMerge {
					output <- testResults
				}
			} else if name == "sample" {
				var sample Sample
				err := decoder.DecodeElement(&sample, &elementType)
				if err != nil {
					fmt.Println("Could not decode element", err)
					// log this, but just skip
					break
				}

				// send sample
				output <- sample
			} else if name == "httpSample" {
				var sample HttpSample
				err := decoder.DecodeElement(&sample, &elementType)
				if err != nil {
					fmt.Println("Could not decode element", err)
					// log this, but just skip
					break
				}

				// send sample
				output <- sample
			} else {
				fmt.Println("Found an unknown element", name)
			}
		default:
		}
	}
}
