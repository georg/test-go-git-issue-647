package server

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestClone(t *testing.T) {
	gitDir, err := filepath.Abs(filepath.Join("..", ".git"))
	if err != nil {
		t.Fatal(err)
	}
	go Start(gitDir)

	// git automatically adds --no-progress with stderr is not a TTY
	cmd := exec.CommandContext(t.Context(), "git", "clone", "http://localhost:8080/")
	cmd.Env = append(cmd.Env, "GIT_TRACE=true")
	cmd.Env = append(cmd.Env, "GIT_TRACE_PACKET=true")
	cmd.Env = append(cmd.Env, "PATH="+os.Getenv("PATH"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = t.TempDir()

	if err = cmd.Run(); err != nil {
		t.Error(err)
	}
}
