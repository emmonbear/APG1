// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/xargs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./myXargs [options] [arguments...]")
		os.Exit(1)
	}

	// myXargs := xargs.New(os.Args[1], os.Args[2:])

	for i, arg := range os.Args {
		
	} 
	fmt.Printf("")
	
	// if err := myXargs.ParseCommandLine(os.Stdin); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error during input processing: %v\n", err)
	// 	os.Exit(1)
	// }

	// if err := myXargs.Execute(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "A command execution error: %v\n", err)
	// 	os.Exit(1)
	// }
}
