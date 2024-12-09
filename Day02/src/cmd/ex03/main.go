// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/rotate"
)

func main() {
	config, err := rotate.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "flag handling error: %v\n", err)
		os.Exit(1)
	}
	err = rotate.Rotate(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rotate error: %v\n", err)
		os.Exit(1)
	}
}
