package main

import (
	"main/server"
	"path/filepath"
)

func main() {
	gitDir, err := filepath.Abs(filepath.Join(".git"))
	if err != nil {
		panic(err)
	}
	server.Start(gitDir)
}
