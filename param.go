package blast

import (
	"fmt"
	"os/exec"
)

const (
	D_TOOL    string  = "blastp"
	D_EVALUE  float64 = 0.001
	D_THREADS int     = 2
	D_OUTFMT  string  = "5"
	D_OUTPUT  string  = "stdout"
)

type Param struct {
	tool    string
	evalue  float64
	threads int
	outfmt  string
	output  string
}

func NewParam() *Param {
	var p Param
	p.tool = D_TOOL
	p.evalue = D_EVALUE
	p.threads = D_THREADS
	p.outfmt = D_OUTFMT
	p.output = D_OUTPUT

	return &p
}

// SETTER
func (p *Param) SetTool(t string) {
	p.tool = t
}
func (p *Param) SetEvalue(e float64) {
	p.evalue = e
}
func (p *Param) SetThreads(t int) {
	p.threads = t
}
func (p *Param) SetOutfmt(o string) {
	p.outfmt = o
}
func (p *Param) SetOutput(o string) {
	p.output = o
}

// GETTER
func (p *Param) GetTool() string {
	return p.tool
}
func (p *Param) GetEvalue() float64 {
	return p.evalue
}
func (p *Param) GetThreads() int {
	return p.threads
}
func (p *Param) GetOutfmt() string {
	return p.outfmt
}
func (p *Param) GetOutput() string {
	return p.output
}

// GENERATE THE COMMAND LINE
func (p *Param) GetCmd(db string) *exec.Cmd {
	if p.output == "stdout" {
		return exec.Command(
			p.tool,
			"-db", db,
			"-evalue", fmt.Sprintf("%f", p.evalue),
			"-num_threads", fmt.Sprintf("%d", p.threads),
			"-outfmt", p.outfmt,
		)
	} else {
		return exec.Command(
			p.tool,
			"-db", db,
			"-evalue", fmt.Sprintf("%f", p.evalue),
			"-num_threads", fmt.Sprintf("%d", p.threads),
			"-outfmt", p.outfmt,
			"-out", p.output,
		)
	}
}
