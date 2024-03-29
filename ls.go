package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

func main() {
	directory, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	contents, err := ioutil.ReadDir(directory)
	if err != nil {
		os.Exit(1)
	}

	writer := tabwriter.NewWriter(color.Output, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)
	printBlue := color.New(color.FgHiBlue).SprintFunc()
	printGreen := color.New(color.FgHiGreen).SprintFunc()

	var stringBuilder strings.Builder

	for _, entry := range contents {

		strings.TrimSuffix(entry.Name(), ".exe")

		if entry.IsDir() {
			stringBuilder.WriteString(printBlue(entry.Name()))
			stringBuilder.WriteString("/\t\n")
		} else if strings.HasSuffix(entry.Name(), ".exe") {
			stringBuilder.WriteString(printGreen(strings.TrimSuffix(entry.Name(), ".exe")))
			stringBuilder.WriteString("\t\n")
		} else {
			stringBuilder.WriteString(entry.Name())
			stringBuilder.WriteString("\t\n")
		}
		/*
			if x%4 == 0 && x > 0 {
				stringBuilder.WriteRune('\n')
			}
		*/
	}

	fmt.Fprint(writer, stringBuilder.String())
	writer.Flush()

	os.Exit(0)
}
