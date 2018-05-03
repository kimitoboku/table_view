package main

import (
	"bufio"
	"flag"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"os"
	"strings"
)

var (
	inputTSV  = flag.Bool("t", false, "read tsv")
	delimitre = ","
)

func main() {
	flag.Parse()
	if *inputTSV {
		delimitre = "\t"
	}

	fn := flag.Arg(0)
	f, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	table := tablewriter.NewWriter(os.Stdout)
	i := 0
	reader := bufio.NewReaderSize(f, 1000)
	for {
		lineB, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		line := string(lineB)
		d := strings.Split(line, delimitre)
		if i == 0 {
			table.SetHeader(d)
		} else {
			table.Append(d)
		}
		i += 1
	}
	table.Render()
}
