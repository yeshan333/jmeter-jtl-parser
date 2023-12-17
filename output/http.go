package output

import (
	"encoding/json"
	"fmt"

	"github.com/yeshan333/jmeter-jtl-parser/jtl"
)

func HTTP(output <-chan interface{}) string {
	var result string
	result += "{"

	first := true
	for {
		element, more := <-output
		if !more {
			result += "]}}"
			break
		}
		switch elementType := element.(type) {
		case jtl.TestResults:
			str := fmt.Sprintf(`"testResults": { "version": "%v", "samples": [`, elementType.Version)
			result += str
		default:
			parsed, err := json.MarshalIndent(element, "", "    ")
			if err != nil {
				fmt.Printf("Could not marshal xml: %v\n", err)
			}
			if !first {
				result += ","
			} else {
				first = false
			}
			result += string(parsed)
		}
	}

	return result
}
