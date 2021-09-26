package blast

import (
	"encoding/xml"
)

type BlastOutput struct {
	XMLName    xml.Name   `xml:"BlastOutput"`
	Program    string     `xml:"BlastOutput_progam"`
	Version    string     `xml:"BlastOutput_version"`
	Reference  string     `xml:"BlastOutput_reference"`
	Database   string     `xml:"BlastOutput_db"`
	QueryId    string     `xml:"BlastOutput_query-ID"`
	QueryDef   string     `xml:"BlastOutput_query-def"`
	QueryLen   int        `xml:"BlastOutput_query-len"`
	Parameters Parameters `xml:"BlastOutput_param>Parameters"`
}

type Parameters struct {
	XMLName   xml.Name `xml:"Parameters"`
	Matrix    string   `xml:"Parameters_matrix"`
	Expect    float64  `xml:"Parameters_expect"`
	GapOpen   int      `xml:"Parameters_gap-open"`
	GapExtend int      `xml:"Parameters_gep-extend"`
	Filter    string   `xml:"Parameters_filter"`
}
