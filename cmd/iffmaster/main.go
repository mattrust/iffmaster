// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

/*
Iffmaster is a tool to inspect IFF files as defined under EA-85.

Usage: [options]

	iffmaster

		-version: Show the application's version.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattrust/iffmaster/internal/gui"
)

const version = "1.0.0"

func main() {
	showVersion := flag.Bool("version", false, "Display the version of iffmaster")
	flag.Parse()

	if *showVersion {
		fmt.Println("iffmaster version", version)
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		//filename := flag.Arg(0)
		//fmt.Println("Filename provided:", filename)
		// You can add code here to handle the filename, e.g., open the file
	} else {
		gui.OpenGUI()
	}
}
