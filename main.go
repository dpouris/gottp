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
	helpSubFlag := getFlag.Bool("h", true, "Learn about the commands")

	// POST argument that makes a post request to a url with a payload specified
	postFlag := flag.NewFlagSet("post", flag.ExitOnError)
	postUrlSubFlag := postFlag.String("u", "", "Specify the url to post data to")
	postPayloadFlag := postFlag.String("p", "", "Specify a path to a JSON/XML file to use as a payload for the request")
	postFilenameSubFlag := postFlag.String("f", "post_file.txt", "Specify the path to save the response")
	postHelpSubFlag := postFlag.Bool("h", true, "Learn about the commands")

	// If the Args are less than 2 which means that only the program was run with no arguments we want to provide the error print the available commands and exit
	if len(os.Args) < 2 {
		fmt.Println("Expected 1 or more arguments, provided none.")
		getFlag.PrintDefaults()
		os.Exit(1)
	}

	// Parse the args provided in the terminal by accessing the os.Args

	// Since the fist index -> 0 on the os.Args is the program we switch the flag provided that is stored the second index -> 1
	switch os.Args[1] {
	case "get":
		getFlag.Parse(os.Args[2:])
		handleGet(getFlag, urlSubFlag, filenameSubFlag, helpSubFlag)
	case "post":
		postFlag.Parse(os.Args[2:])
		handlePost(postFlag, postPayloadFlag, postFilenameSubFlag, postUrlSubFlag, postHelpSubFlag)
	}

}

func handlePost(pf *flag.FlagSet, p *string, f *string, u *string, h *bool) {
	if !(*h) {
		pf.PrintDefaults()
		os.Exit(0)
		return
	}

	if *u == "" {
		fmt.Println("Please provide a url to post to.")
		pf.PrintDefaults()
		os.Exit(1)
		return
	}

	resp, err := postUrl(*u, p)

	if err != nil {
		fmt.Println("Something went wrong with the request. Please make sure you're providing a valid url and or payload")
		os.Exit(1)
		return
	}

	fw := fileWriter{
		fileDst: *f,
	}

	if *f != "" {
		byteStream := make([]byte, 10*1024)
		resp.Body.Read(byteStream)
		resp.Body.Close()
		fw.Write(byteStream)
	}

	fmt.Print("\n-------------------- RESPONSE --------------------\n\n")
	fmt.Println(resp)
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
