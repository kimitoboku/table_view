package main

import (
	"bufio"
	"flag"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"os"
	"fmt"
	"strings"
)

var (
	inputTSV  = flag.Bool("t", false, "read tsv")
	outputLaTeX = flag.Bool("l", false, "Output LaTeX Table")
	delimitre = ","
)

func printCLI(fn string) {
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
		i++
	}
	table.Render()
	return
}

func printTex(fn string){
	f, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

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
			fmt.Print("\\begin{tabular}{")
			fmt.Print(strings.Repeat("l", len(d)))
			fmt.Println("}")
			fmt.Println(strings.Join(d, " & ") + "\\\\")
			fmt.Println("\\hline")
		} else{
			fmt.Println(strings.Join(d, " & ") + "\\\\")
		}
		i++
	}
	fmt.Println("\\end{tabular}")
	return
}

func main() {
	flag.Parse()
	if *inputTSV {
		delimitre = "\t"
	}

	fn := flag.Arg(0)
	if *outputLaTeX {
		printTex(fn)
	}else{
		printCLI(fn)
	}
}
