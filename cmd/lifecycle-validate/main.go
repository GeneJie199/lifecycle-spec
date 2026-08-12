package main

import (
	"fmt"
	"os"

	lifecyclespec "github.com/GeneJie199/lifecycle-spec"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "--list" {
		names, err := lifecyclespec.SchemaNames()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		for _, name := range names {
			fmt.Println(name)
		}
		return
	}
	var schema, document string
	if len(args) == 2 {
		schema, document = args[0], args[1]
	} else if len(args) == 3 && args[0] == "--schema" {
		schema, document = args[1], args[2]
	} else {
		fmt.Fprintln(os.Stderr, "usage: lifecycle-validate [--schema] <schema-name> <document.json>\n       lifecycle-validate --list")
		os.Exit(2)
	}
	b, err := os.ReadFile(document)
	if err == nil {
		err = lifecyclespec.Validate(schema, b)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid:", err)
		os.Exit(1)
	}
	fmt.Printf("valid: %s against %s\n", document, schema)
}
