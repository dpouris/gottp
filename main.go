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
	postFilenameSubFlag := postFlag.String("f", "", "Specify the path to save the response")
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
	// If the -h flag is specified print the usage of each flag and exit with success code 0
	if !(*h) {
		pf.PrintDefaults()
		os.Exit(0)
		return
	}

	// If the dereferenced flag u is a zero value then give reminder, print defaults and exit with code 1
	if *u == "" {
		fmt.Println("Please provide a url to post to.")
		pf.PrintDefaults()
		os.Exit(1)
		return
	}

	// Post to the url specified with the p which is a string representing the path to the file in which lays the paylaod
	resp, err := postUrl(*u, p)

	//Handle request error
	if err != nil {
		fmt.Println("Something went wrong with the request. Please make sure you're providing a valid url and or payload")
		os.Exit(1)
		return
	}

	// Read the body into []byte
	byteStream := make([]byte, 30*1024)
	resp.Body.Read(byteStream)
	// Remove all the unused space in the slice
	byteStream = zeroByteStripper(byteStream)
	resp.Body.Close()

	// instantiate a new fileWriter in order to write to a file if specified by the flag
	if *f != "" {
		fw := fileWriter{
			fileDst: *f,
		}
		fw.Write(byteStream)
	}

	fmt.Print("\n-------------------- RESPONSE --------------------\n\n")
	fmt.Println(resp)
	fmt.Print("\n-------------------- BODY --------------------\n\n")
	fmt.Println(string(byteStream))
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
