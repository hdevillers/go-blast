package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hdevillers/go-blast"
)

func main() {
	f, e := os.Open("test.xml")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	x, e := ioutil.ReadAll(f)

	var bo blast.BlastOutput

	xml.Unmarshal(x, &bo)

	fmt.Println(bo)
}
