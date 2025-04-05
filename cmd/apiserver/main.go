// Package main is the entry point for the application.
package main

import "github.com/minuk-dev/minuk-boilerplate/pkg/cmd/apiserver"

func main() {
	cmd := apiserver.NewCommand(apiserver.Options{})
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
