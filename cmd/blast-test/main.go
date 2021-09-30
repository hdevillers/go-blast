package main

import (
	"flag"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seqio"
)

func main() {
	query := flag.String("query", "", "Input query file.")
	format := flag.String("format", "fasta", "Input query format.")
	db := flag.String("db", "", "BLAST Database path.")
	header := flag.String("header", "", "Required header for the summary.")
	maxHit := flag.Int("max-hits", 100, "Maximal number of hit to report.")

	flag.Parse()

	if *query == "" {
		panic("You must provide an input query file.")
	}

	if *db == "" {
		panic("You must provide an BLASTable database.")
	}

	// Init blast object
	blast := blast.NewBlast()
	blast.Db = *db

	// Load query files
	reader := seqio.NewReader(*query, *format, false)
	reader.CheckPanic()
	defer reader.Close()
	for reader.Next() {
		reader.CheckPanic()
		blast.AddQuery(reader.Seq())
	}

	// Launch the blast search
	err := blast.Search()
	if err != nil {
		panic(err)
	}

	blast.Rst.BestHspSummary(*header, *maxHit)
}
