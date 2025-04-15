package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"os/exec"
	"strings"
)

const (
	maxLength = 35
	baseDir   = "/var/www/werc/sites/c.bauherren.ovh"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide one search term")
		os.Exit(1)
	}

	searchTerm := os.Args[1]
	if len(searchTerm) > maxLength {
		fmt.Printf("Search term is too long. Maximum allowed length is %d characters.\n", maxLength)
		os.Exit(1)
	}

	fmt.Printf("<link rel=stylesheet href=/_werc/pub/search-results.css><div id=\"search-results\">\n")
	fmt.Printf("<h3>Search results for: \"%s\"</h3>\n", html.EscapeString(searchTerm))

	cmd := exec.Command("rg", "-i",  "--no-heading", "--with-filename", "--line-number", searchTerm, "--glob", "*.md", baseDir)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("No results")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}

		file := strings.TrimPrefix(parts[0], baseDir)
		file = strings.ReplaceAll(file, ".md", "")
		lineNumber := parts[1]
		content := parts[2]

		escapedContent := html.EscapeString(content)
		formattedContent := strings.ReplaceAll(escapedContent, searchTerm, "<b>"+searchTerm+"</b>")

		fmt.Printf("<div class=\"search-item\">\n")
		fmt.Printf("<a href=\"%s\">%s</a> (Line %s)<br>\n", file, file, lineNumber)
		fmt.Printf("<p>\"%s\"</p>\n", formattedContent)
		fmt.Printf("</div>\n")
	}

	fmt.Printf("</div>\n")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading output:", err)
		os.Exit(1)
	}
}

