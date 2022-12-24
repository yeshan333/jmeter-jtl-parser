package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/yeshan333/jmeter-jtl-parser/jtl"
	"github.com/yeshan333/jmeter-jtl-parser/output"
)

const DefaultOutputFormat = "http"

type OutputType string

const (
	CSV  OutputType = "csv"
	XML  OutputType = "xml"
	JSON OutputType = "json"
	HTTP OutputType = "http"
)

func main() {
	var args struct {
		Filename      string   `arg:"positional,required"`
		Output        string   `arg:"-o" default:"http" help:"specify the output type, valid options: csv,xml,json,http"`
		MergedFiles   []string `arg:"-f" help:"merge jtl files"`
		ReportApiHost string   `arg:"-h" default:"http://localhost:8080/jmeter"`
	}
	args.Output = DefaultOutputFormat
	arg.MustParse(&args)

	// type Args struct {
	// 	Filename string
	// 	Output   string
	// }

	// args := &Args{Filename: "data.jtl", Output: "http"}
	outputType := OutputType(args.Output)

	data := make(chan interface{})
	go jtl.Decode(args.Filename, data, args.MergedFiles)

	if outputType == HTTP {
		postData := output.HTTP(data)
		postToReportApi(args.ReportApiHost, postData)
	} else {
		outputStream := getOutputStream(outputType)
		outputStream(data)
	}
}

func getOutputStream(outputType OutputType) func(data <-chan interface{}) {
	switch outputType {
	case CSV:
		return output.CSV
	case XML:
		return output.XML
	case JSON:
		return output.JSON
	default:
		panic("Unknown output type " + outputType)
	}
}

// post data to test report api
func postToReportApi(jmeterReportHost string, postData string) {
	url := jmeterReportHost
	method := "POST"

	payload := strings.NewReader(postData)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
