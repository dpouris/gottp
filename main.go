package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// GET argument that fetches a from a -u or --url
	getFlag := flag.NewFlagSet("get", flag.ExitOnError)
	urlSubFlag := getFlag.String("u", "", "Specify the url to fetch data from")
	filenameSubFlag := getFlag.String("f", "file.txt", "Specify the path to save the fetched data")
	helpSubFlag := getFlag.Bool("h", true, "Learn abou the commands")

	// If the Args are less than 2 which means that only the program was run with no arguments we want to provide the error print the available commands and exit
	if len(os.Args) < 2 {
		fmt.Println("Expected 1 or more arguments, provided none.")
		getFlag.PrintDefaults()
		os.Exit(1)
	}

	// Parse the args provided in the terminal by accessing the os.Args
	getFlag.Parse(os.Args[2:])

	// Since the fist index -> 0 on the os.Args is the program we switch the flag provided that is stored the second index -> 1
	switch os.Args[1] {
	case "get":
		handleGet(getFlag, urlSubFlag, filenameSubFlag, helpSubFlag)
	}

}

// The get flag || arg handler
func handleGet(g *flag.FlagSet, u *string, f *string, h *bool) {
	if !(*h) {
		g.PrintDefaults()
		os.Exit(0)
		return
	}

	if *u == "" {
		fmt.Println("Must provide a url to fetch")
		g.PrintDefaults()
		os.Exit(1)
	}

	dataStream, err := getUrl(*u)

	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}

	fw := fileWriter{
		fileDst: *f,
	}

	fw.Write(dataStream)

	fmt.Print("\n-------------------- DATA --------------------\n\n")
	fmt.Println(string(dataStream))
}
