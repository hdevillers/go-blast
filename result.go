package blast

import (
	"encoding/xml"
)

type BlastOutput struct {
	XMLName    xml.Name    `xml:"BlastOutput"`
	Program    string      `xml:"BlastOutput_progam"`
	Version    string      `xml:"BlastOutput_version"`
	Reference  string      `xml:"BlastOutput_reference"`
	Database   string      `xml:"BlastOutput_db"`
	QueryId    string      `xml:"BlastOutput_query-ID"`
	QueryDef   string      `xml:"BlastOutput_query-def"`
	QueryLen   int         `xml:"BlastOutput_query-len"`
	Parameters Parameters  `xml:"BlastOutput_param>Parameters"`
	Iterations []Iteration `xml:"BlastOutput_iterations>Iteration"`
}

type Parameters struct {
	XMLName   xml.Name `xml:"Parameters"`
	Matrix    string   `xml:"Parameters_matrix"`
	Expect    float64  `xml:"Parameters_expect"`
	GapOpen   int      `xml:"Parameters_gap-open"`
	GapExtend int      `xml:"Parameters_gep-extend"`
	Filter    string   `xml:"Parameters_filter"`
}

type Iteration struct {
	XMLName  xml.Name `xml:"Iteration"`
	IterNum  int      `xml:"Iteration_iter-num"`
	QueryId  string   `xml:"Iteration_query-ID"`
	QueryDef string   `xml:"Iteration_query-def"`
	QueryLen int      `xml:"Iteration_query-len"`
	Hits     []Hit    `xml:"Iteration_hits>Hit"`
}

type Hit struct {
	XMLName      xml.Name `xml:"Hit"`
	HitNum       int      `xml:"Hit_num"`
	HitId        string   `xml:"Hit_id"`
	HitDef       string   `xml:"Hit_def"`
	HitAccession int      `xml:"Hit_accession"`
	HitLen       int      `xml:"Hit_len"`
	HitHsps      []Hsp    `xml:"Hit_hsps>Hsp"`
}

type Hsp struct {
	XMLName    xml.Name `xml:"Hsp"`
	BitScore   float64  `xml:"Hsp_bit-score"`
	Score      int      `xml:"Hsp_score"`
	Evalue     float64  `xml:"Hsp_evalue"`
	QueryFrom  int      `xml:"Hsp_query-from"`
	QueryTo    int      `xml:"Hsp_query-to"`
	QueryFrame int      `xml:"Hsp_query-frame"`
	HitFrom    int      `xml:"Hsp_hit-from"`
	HitTo      int      `xml:"Hsp_hit-to"`
	HitFrame   int      `xml:"Hsp_hit-frame"`
	Identity   int      `xml:"Hsp_identity"`
	Similarity int      `xml:"Hsp_positive"`
	Gaps       int      `xml:"Hsp_gaps"`
	AlignLen   int      `xml:"Hsp_align-len"`
	QuerySeq   string   `xml:"Hsp_qseq"`
	HitSeq     string   `xml:"Hsp_hseq"`
	AlignSeq   string   `xml:"Hsp_midline"`
}

func NewBlastOutput() *BlastOutput {
	return &BlastOutput{}
}
