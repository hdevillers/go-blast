package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seqio"
)

func main() {
	query := flag.String("query", "", "Input query file.")
	format := flag.String("format", "fasta", "Input query format.")
	db := flag.String("db", "", "BLAST database path or FASTA database.")
	header := flag.String("header", "", "Required header for the summary.")
	maxHit := flag.Int("max-hits", 100, "Maximal number of hit to report.")
	tool := flag.String("tool", "blastp", "BLAST+ tool to use.")
	task := flag.String("task", "", "Specify the BLAST+ task to run (default = same as tool).")
	evalue := flag.Float64("evalue", 0.001, "E-value threshold.")
	filterLC := flag.Bool("filter-lc", false, "Filter low complexity region in query.")
	threads := flag.Int("threads", 2, "Number of threads.")
	showTool := flag.Bool("tool-list", false, "Display the list of available tools.")

	flag.Parse()

	allowedTools := map[string]string{}
	allowedTools["blastn"] = "Nucl => Nucl comparison."
	allowedTools["blastp"] = "Prot => Prot comparison."
	allowedTools["tblastn"] = "Prot=>Trans(Nucl) comparison."
	allowedTools["blastx"] = "Trans(Nucl)=>Prot comparison."
	allowedTools["tblastx"] = "Trans(Nucl)=>Trans(Nucl) comparison."

	if *showTool {
		for k, v := range allowedTools {
			fmt.Printf("%s\t%s\n", k, v)
		}
		os.Exit(1)
	}

	// Check tool value
	if _, ok := allowedTools[*tool]; !ok {
		panic(fmt.Sprintf("The tool %s does not exist!", *tool))
	}

	if *query == "" {
		panic("You must provide an input query file.")
	}

	if *db == "" {
		panic("You must provide an BLASTable database.")
	}

	// Init blast object
	blast := blast.NewBlast()
	blast.Db = *db

	// Setup argument values
	blast.Par.SetTool(*tool)
	if *task == "" {
		// Default task is the tool
		*task = *tool
	}
	blast.Par.SetTask(*task)
	blast.Par.SetFilterLC(*filterLC)
	blast.Par.SetEvalue(*evalue)
	blast.Par.SetThreads(*threads)

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
