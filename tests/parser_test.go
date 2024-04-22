package main

import (
	"testing"

	"github.com/yeshan333/jmeter-jtl-parser/jtl"
	"github.com/yeshan333/jmeter-jtl-parser/output"
)

func TestXMLOutputToCSV(t *testing.T) {
	data := make(chan interface{})
	go jtl.Decode("./github-etcd.jtl", data, []string{})
	output.CSV(data)
}

func TestXMLOutputToJSON(t *testing.T) {
	data := make(chan interface{})
	go jtl.Decode("./github-etcd.jtl", data, []string{})
	output.JSON(data)
}
