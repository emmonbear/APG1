package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "1",
			args: []string{"main", "-f", "./"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := os.Args
		
			defer func() {
				os.Args = args
			}()
		
			os.Args = tt.args
			main()

		})
	}


}