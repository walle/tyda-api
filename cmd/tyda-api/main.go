package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/walle/tyda-api"
)

var args struct {
	Query     string   `arg:"positional,required"`
	Indented  bool     `arg:"-i,env,help:If the output should be indented"`
	Languages []string `arg:"-l,env,help:Languages to translate to (en fr de es la nb sv)"`
}

func main() {
	args.Languages = []string{"en"}
	arg.MustParse(&args)

	response, err := tydaapi.Search(args.Query, args.Languages)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong when searching: %s\n", err)
		os.Exit(1)
	}

	if args.Indented {
		buffer, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Something went wrong printing result: %s\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s\n", buffer)
	} else {
		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Something went wrong printing result: %s\n", err)
			os.Exit(1)
		}
	}
}
