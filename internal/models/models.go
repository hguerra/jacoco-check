package models

import "encoding/xml"

type SessionInfo struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id,attr"`
	Start string `xml:"start,attr"`
	Dump  string `xml:"dump,attr"`
}

type Line struct {
	Text string `xml:",chardata"`
	Nr   string `xml:"nr,attr"`
	Mi   string `xml:"mi,attr"`
	Ci   string `xml:"ci,attr"`
	Mb   string `xml:"mb,attr"`
	Cb   string `xml:"cb,attr"`
}

type Counter struct {
	Text    string `xml:",chardata"`
	Type    string `xml:"type,attr"`
	Missed  string `xml:"missed,attr"`
	Covered string `xml:"covered,attr"`
}

type SourceFiles struct {
	Text     string    `xml:",chardata"`
	Name     string    `xml:"name,attr"`
	Lines    []Line    `xml:"line"`
	Counters []Counter `xml:"counter"`
}

type Method struct {
	Text     string    `xml:",chardata"`
	Name     string    `xml:"name,attr"`
	Desc     string    `xml:"desc,attr"`
	Line     string    `xml:"line,attr"`
	Counters []Counter `xml:"counter"`
}

type Class struct {
	Text           string    `xml:",chardata"`
	Name           string    `xml:"name,attr"`
	SourceFilename string    `xml:"sourcefilename,attr"`
	Methods        []Method  `xml:"method"`
	Counters       []Counter `xml:"counter"`
}

type Package struct {
	Text        string        `xml:",chardata"`
	Name        string        `xml:"name,attr"`
	Classes     []Class       `xml:"class"`
	SourceFiles []SourceFiles `xml:"sourcefile"`
	Counters    []Counter     `xml:"counter"`
}

type Report struct {
	XMLName     xml.Name    `xml:"report"`
	Text        string      `xml:",chardata"`
	Name        string      `xml:"name,attr"`
	SessionInfo SessionInfo `xml:"sessioninfo"`
	Packages    []Package   `xml:"package"`
	Counters    []Counter   `xml:"counter"`
}
