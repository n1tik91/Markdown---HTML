package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/russross/blackfriday/v2"
)

func main() {
	// Check if the user provided a file name
	if len(os.Args) < 2 {
		fmt.Println("Please provide a markdown file to convert.")
		return
	}

	// Get the markdown file name from command-line arguments
	markdownFile := os.Args[1]

	// Read the content of the markdown file
	data, err := ioutil.ReadFile(markdownFile)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Convert the markdown content to HTML
	htmlContent := blackfriday.Run(data)

	// Create a new HTML file name
	htmlFileName := markdownFile + ".html"

	// Write the HTML content to a new file
	err = ioutil.WriteFile(htmlFileName, htmlContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Inform the user the conversion is complete
	fmt.Println("Conversion complete. HTML saved as:", htmlFileName)
}