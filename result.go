package blast

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	HitAccession string   `xml:"Hit_accession"`
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

/* COMPUTE HSP/HIT STAT */
const (
	HEV string = "evalue"
	HBS string = "bit.score"

	HQC string = "query.cover"
	HQF string = "query.from"
	HQG string = "query.gap"
	HQI string = "query.identity"
	HQL string = "query.length"
	HQP string = "query.strand"
	HQR string = "query.seq"
	HQS string = "query.similarity"
	HQT string = "query.to"

	HHC string = "hit.cover"
	HHG string = "hit.gap"
	HHF string = "hit.from"
	HHI string = "hit.identity"
	HHL string = "hit.length"
	HHP string = "hit.strand"
	HHR string = "hit.seq"
	HHS string = "hit.similarity"
	HHT string = "hit.to"

	HAG string = "align.gap"
	HAI string = "align.identity"
	HAL string = "align.length"
	HAR string = "align.seq"
	HAS string = "align.similarity"

	HRG string = "raw.gaps"
	HRI string = "raw.identity"
	HSC string = "raw.score"
	HRS string = "raw.similarity"

	DRO string = "%03f"
)

func (h *Hsp) GetHspStat(ql int, hl int) map[string]float64 {
	hs := make(map[string]float64)

	// Check query and hit coordinate
	qf := h.QueryFrom
	qt := h.QueryTo
	if qt < qf {
		qf, qt = qt, qf
	}
	hf := h.HitFrom
	ht := h.HitTo
	if qt < qf {
		hf, ht = ht, hf
	}
	hs[HQC] = float64(qt-qf+1) / float64(ql) * 100.0
	hs[HQS] = float64(h.Similarity) / float64(ql) * 100.0
	hs[HQI] = float64(h.Identity) / float64(ql) * 100.0
	hs[HQG] = float64(h.Gaps) / float64(ql) * 100.0

	hs[HHC] = float64(ht-hf+1) / float64(hl) * 100.0
	hs[HHS] = float64(h.Similarity) / float64(hl) * 100.0
	hs[HHI] = float64(h.Identity) / float64(hl) * 100.0
	hs[HHG] = float64(h.Gaps) / float64(hl) * 100.0

	hs[HAS] = float64(h.Similarity) / float64(h.AlignLen) * 100.0
	hs[HAI] = float64(h.Identity) / float64(h.AlignLen) * 100.0
	hs[HAG] = float64(h.Gaps) / float64(h.AlignLen) * 100.0

	return hs
}

func (h *Hsp) GetHspDetails(ql, hl int) map[string]string {
	hd := make(map[string]string)

	// Compute stat
	hs := h.GetHspStat(ql, hl)

	// Create the complete data map
	hd[HEV] = fmt.Sprintf("%03e", h.Evalue)
	hd[HSC] = fmt.Sprintf("%d", h.Score)
	hd[HBS] = fmt.Sprintf(DRO, h.BitScore)

	hd[HQC] = fmt.Sprintf(DRO, hs[HQC])
	hd[HQS] = fmt.Sprintf(DRO, hs[HQS])
	hd[HQI] = fmt.Sprintf(DRO, hs[HQI])
	hd[HQG] = fmt.Sprintf(DRO, hs[HQG])
	hd[HQL] = fmt.Sprintf("%d", ql)
	hd[HQP] = fmt.Sprintf("%d", h.QueryFrame)
	hd[HQR] = h.QuerySeq
	hd[HQF] = fmt.Sprintf("%d", h.QueryFrom)
	hd[HQT] = fmt.Sprintf("%d", h.QueryTo)

	hd[HHC] = fmt.Sprintf(DRO, hs[HHC])
	hd[HHS] = fmt.Sprintf(DRO, hs[HHS])
	hd[HHI] = fmt.Sprintf(DRO, hs[HHI])
	hd[HHG] = fmt.Sprintf(DRO, hs[HHG])
	hd[HHL] = fmt.Sprintf("%d", hl)
	hd[HHP] = fmt.Sprintf("%d", h.HitFrame)
	hd[HHR] = h.HitSeq
	hd[HHF] = fmt.Sprintf("%d", h.HitFrom)
	hd[HHT] = fmt.Sprintf("%d", h.HitTo)

	hd[HAS] = fmt.Sprintf(DRO, hs[HAS])
	hd[HAI] = fmt.Sprintf(DRO, hs[HAI])
	hd[HAG] = fmt.Sprintf(DRO, hs[HAG])
	hd[HAL] = fmt.Sprintf("%d", h.AlignLen)
	hd[HAR] = h.AlignSeq

	hd[HRG] = fmt.Sprintf("%d", h.Gaps)
	hd[HRI] = fmt.Sprintf("%d", h.Identity)
	hd[HRS] = fmt.Sprintf("%d", h.Similarity)

	return hd
}

type HeaderSummary struct {
	RequireHeaders []string
	StillHeaders   []string
	DefaultHeaders []string
	AllowedHeaders map[string]int
}

func NewHeaderSummary() *HeaderSummary {
	var hs HeaderSummary

	// List of all header allowed
	hs.AllowedHeaders = map[string]int{
		HEV: 1, HBS: 1, HQC: 1, HQF: 1, HQG: 1, HQI: 1,
		HQL: 1, HQP: 1, HQR: 1, HQS: 1, HQT: 1, HHC: 1,
		HHG: 1, HHF: 1, HHI: 1, HHL: 1, HHP: 1, HHR: 1,
		HHS: 1, HHT: 1, HAG: 1, HAI: 1, HAL: 1, HAR: 1,
		HAS: 1, HRG: 1, HRI: 1, HSC: 1, HRS: 1,
	}

	// Still header
	hs.StillHeaders = []string{"query.id", "hit.id", "hit.num"}

	// Default header
	hs.DefaultHeaders = []string{HEV, HBS, HQC, HQS, HQI}

	return &hs
}

// Print only best hsp summary
func (bo *BlastOutput) BestHspSummary(qh string, mh int) {
	if len(bo.Iterations) == 0 {
		panic("Blast output is empty!")
	}

	// Parse the qury header
	hs := NewHeaderSummary()
	if qh != "" {
		for _, he := range strings.Split(qh, ",") {
			if _, ok := hs.AllowedHeaders[he]; ok {
				hs.RequireHeaders = append(hs.RequireHeaders, he)
			} else {
				panic(fmt.Sprintf("The require header key %s does not exists.", he))
			}
		}
	} else {
		// Use default header
		hs.RequireHeaders = hs.DefaultHeaders
	}

	// Print the header
	for _, i := range hs.StillHeaders {
		fmt.Printf("%s\t", i)
	}
	fmt.Printf("hsp.count")
	for _, i := range hs.RequireHeaders {
		fmt.Printf("\t%s", i)
	}
	fmt.Printf("\n")

	// Scan each queries and hits
	for _, iter := range bo.Iterations {
		if len(iter.Hits) == 0 {
			// No hits
			fmt.Printf("%s\tNA\tNA\tNA", iter.QueryDef)
			for range hs.RequireHeaders {
				fmt.Printf("\tNA")
			}
			fmt.Printf("\n")
		} else {
			nhit := 0
		HITS:
			for _, hit := range iter.Hits {
				hsp := hit.HitHsps[0]
				fmt.Printf("%s\t%s\t%d\t%d", iter.QueryDef, hit.HitDef, hit.HitNum, len(hit.HitHsps))
				hd := hsp.GetHspDetails(iter.QueryLen, hit.HitLen)
				for _, he := range hs.RequireHeaders {
					fmt.Printf("\t%s", hd[he])
				}
				fmt.Printf("\n")
				nhit++
				if nhit == mh {
					break HITS
				}
			}
		}
	}
}
