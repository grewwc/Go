package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/grewwc/go_tools/src/terminalW"
	"github.com/grewwc/go_tools/src/utilsW"
)

func init() {
	terminalW.EnableVirtualTerminal()
}

func main() {
	var numOfLines = 10
	parsedResults := terminalW.ParseArgsCmd()
	if parsedResults == nil {
		return
	}

	filenames := parsedResults.Positional.ToStringSlice()

	if nStr, exists := parsedResults.Optional["-n"]; exists {
		// delete(parsedResults.Optional, "-n")
		if nStr == "" {
			return
		}
		n, err := strconv.ParseInt(nStr, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		numOfLines = int(n)
	}

	n := parsedResults.GetNumArgs()
	if n != -1 {
		numOfLines = n
	}

	for _, filename := range filenames {
		if utilsW.IsDir(filename) {
			continue
		}
		f, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(color.HiGreenString("=======>\t%s\n", filename))
		scanner := bufio.NewScanner(f)
		count := 0
		for scanner.Scan() && count < numOfLines {
			line := scanner.Text()
			count++
			fmt.Printf("\t%s\n", line)
		}

		f.Close()

		fmt.Printf("\n\n")
	}
}
