package main

import (
	"flag"
	"fmt"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seqio"
)

func main() {
	query := flag.String("query", "", "Input query file.")
	db := flag.String("db", "", "BLAST Database path.")
	format := flag.String("format", "fasta", "Input query format.")
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

	//blast.Par.SetOutput("test.blast")

	// Load query files
	reader := seqio.NewReader(*query, *format, false)
	reader.CheckPanic()
	defer reader.Close()

	for reader.Next() {
		reader.CheckPanic()
		blast.AddQuery(reader.Seq())
	}

	err := blast.Search()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(blast.Rst.Iterations[0].Hits))
}
