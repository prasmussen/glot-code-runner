package kotlin

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
	"strings"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	fname := filepath.Base(files[0])

	stdout, stderr, err := cmd.Run(workDir, "kotlinc", fname)
	if err != nil {
		return stdout, stderr, err
	}

	return cmd.RunStdin(workDir, stdin, "kotlin", className(fname))
}

func className(fname string) string {
	if len(fname) < 5 {
		return fname
	}

	ext := filepath.Ext(fname)
	name := fname[0 : len(fname)-len(ext)]
	return strings.ToUpper(string(name[0])) + name[1:] + "Kt"
}
