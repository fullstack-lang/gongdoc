package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: program <sourcefile> <destfile>")
		os.Exit(1)
	}

	sourcefile := os.Args[1]
	destfile := os.Args[2]

	// Read the source file
	data, err := ioutil.ReadFile(sourcefile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	content := string(data)

	// Create a buffer for the output
	var out bytes.Buffer
	encoder := xml.NewEncoder(&out)

	// Write the start of the SVG tag
	encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "svg", Space: "http://www.w3.org/2000/svg"}})

	// Find SVG tags and extract the content
	for start := strings.Index(content, "<svg"); start != -1; start = strings.Index(content, "<svg") {
		end := strings.Index(content, "</svg>")
		if end == -1 {
			fmt.Println("Unmatched SVG tag.")
			os.Exit(1)
		}

		// Extract the SVG content, excluding the <svg> and </svg> tags
		svgContent := content[start+len("<svg>") : end]

		// update content
		content = content[end+len("</svg>"):]

		// Create a decoder
		decoder := xml.NewDecoder(strings.NewReader(svgContent))

		// Loop over the tokens
		for {
			t, err := decoder.Token()
			if err != nil {
				break
			}

			// Check the token type and write non-comment tokens
			switch t := t.(type) {
			case xml.Comment:
				// Skip comments
				continue
			default:
				if err := encoder.EncodeToken(t); err != nil {
					fmt.Println("Error encoding token:", err)
					os.Exit(1)
				}
			}
		}

		// Flush the encoder
		if err := encoder.Flush(); err != nil {
			fmt.Println("Error flushing encoder:", err)
			os.Exit(1)
		}
	}

	// Write the end of the SVG tag
	encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "svg"}})

	// End the encoder element
	if err := encoder.Flush(); err != nil {
		fmt.Println("Error ending document:", err)
		os.Exit(1)
	}

	// Write the output to the destination file
	if err := ioutil.WriteFile(destfile, out.Bytes(), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}

	fmt.Println("Finished processing", sourcefile)
}
