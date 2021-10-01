package blast

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	D_TOOL     string  = "blastp"
	D_EVALUE   float64 = 0.001
	D_THREADS  int     = 2
	D_OUTFMT   string  = "5"
	D_OUTPUT   string  = "stdout"
	D_FILTERLC bool    = false
)

type Param struct {
	tool     string
	task     string
	filterLC bool
	evalue   float64
	threads  int
	outfmt   string
	output   string
	chkTask  map[string]map[string]int
}

func NewParam() *Param {
	var p Param
	p.tool = D_TOOL
	p.task = D_TOOL // Default task is the tool itself
	p.evalue = D_EVALUE
	p.threads = D_THREADS
	p.outfmt = D_OUTFMT
	p.output = D_OUTPUT
	p.filterLC = D_FILTERLC
	p.chkTask = make(map[string]map[string]int)
	p.chkTask["blastn"] = map[string]int{"blastn": 1, "megablast": 1, "dc-megablast": 1, "blastn-short": 1}
	p.chkTask["blastp"] = map[string]int{"blastp": 1, "blastp-short": 1, "blastp-fast": 1}
	p.chkTask["tblastn"] = map[string]int{"tblastn": 1, "tblastn-fast": 1}
	p.chkTask["blastx"] = map[string]int{"blastx": 1, "blastx-fast": 1}
	return &p
}

// SETTER
func (p *Param) SetTool(t string) {
	p.tool = t
}
func (p *Param) SetTask(t string) {
	p.task = t
}
func (p *Param) SetEvalue(e float64) {
	p.evalue = e
}
func (p *Param) SetFilterLC(f bool) {
	p.filterLC = f
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
func (p *Param) GetTask() string {
	return p.task
}
func (p *Param) GetEvalue() float64 {
	return p.evalue
}
func (p *Param) GetFilterLC() bool {
	return p.filterLC
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

func (p *Param) CheckTask() bool {
	_, test := p.chkTask[p.tool][p.task]
	return test
}

// GENERATE THE COMMAND LINE
func (p *Param) GetCmd(db string) *exec.Cmd {
	cmd := exec.Command(
		p.tool,
		"-evalue", fmt.Sprintf("%f", p.evalue),
		"-num_threads", fmt.Sprintf("%d", p.threads),
		"-outfmt", p.outfmt,
	)
	// If defined a dedicated output file
	if p.output != "stdout" {
		cmd.Args = append(cmd.Args, "-out", p.output)
	}
	// Add the DB of the subject
	// If the file exists, then it must be a subject file
	if _, err := os.Stat(db); os.IsNotExist(err) {
		cmd.Args = append(cmd.Args, "-db", db)
	} else {
		cmd.Args = append(cmd.Args, "-subject", db)
	}
	// No task for tblastx
	if p.tool != "tblastx" {
		// Check task
		if p.CheckTask() {
			cmd.Args = append(cmd.Args, "-task", p.task)
		} else {
			panic(fmt.Sprintf("The task %s is not available for the tool %s.", p.task, p.tool))
		}
	}
	// Manage low complexity filters
	farg := "-seg"
	fval := "no"
	if p.tool == "blastn" {
		farg = "-dust"
	}
	if p.filterLC {
		fval = "yes"
	}
	cmd.Args = append(cmd.Args, farg, fval)

	return cmd
}
