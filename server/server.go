package server

import (
	"fmt"
	"net/http"

	"github.com/go-git/go-billy/v6/osfs"
	githttp "github.com/go-git/go-git/v6/backend/http"
	"github.com/go-git/go-git/v6/plumbing/transport"
)

func Start(gitDir string) {
	fs := osfs.New(gitDir, osfs.WithBoundOS())
	handler := githttp.NewBackend(transport.NewFilesystemLoader(fs, true))

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
