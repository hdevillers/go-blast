package blast

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"io"

	"github.com/hdevillers/go-seq/seq"
	"github.com/hdevillers/go-seq/seqio/fasta"
)

type Blast struct {
	Queries  []seq.Seq
	Nqueries int
	Db       string
	Par      *Param
	Rst      *BlastOutput
}

func NewBlast() *Blast {
	var b Blast
	b.Par = NewParam()
	b.Rst = NewBlastOutput()
	return &b
}

func (b *Blast) AddQuery(s seq.Seq) {
	b.Queries = append(b.Queries, s)
	b.Nqueries++
}

func (b *Blast) ResetQuery() {
	b.Queries = nil
	b.Nqueries = 0
}

func (b *Blast) Search() error {
	// Control if a query has been set up
	if b.Nqueries == 0 {
		return errors.New("No query provided in the Blast object.")
	}

	// Control if a database has been set up
	if b.Db == "" {
		return errors.New("No DB provided in the Blast object.")
	}

	// Get the command line
	cmd := b.Par.GetCmd(b.Db)

	// Throw the query sequence(s) in stdin
	var queryIn bytes.Buffer
	bb := bufio.NewWriter(&queryIn)
	faw := fasta.NewWriter(bb)
	for _, s := range b.Queries {
		faw.Write(s)
		faw.Flush()
	}
	cmd.Stdin = &queryIn

	// Run the command and catch stdout and stderr
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	// Parse the output if possible
	if b.Par.GetOutput() == "stdout" && b.Par.GetOutfmt() == "5" {
		err = xml.Unmarshal(out, b.Rst)

		if err != nil && err != io.EOF {
			return err
		}
	}

	return nil
}
