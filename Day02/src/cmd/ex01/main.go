// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/emmonbear/APG1/Day02/src/pkg/wc"
	"os"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fs := flag.NewFlagSet("main", flag.ContinueOnError)

	options := wc.NewWCFlags()
	err := options.ParseFlags(fs, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%+v", options)
	var wg sync.WaitGroup
	for _, filename := range fs.Args() {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			_, err := wc.WC(file, options)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}(filename)
	}

	wg.Wait()
}
