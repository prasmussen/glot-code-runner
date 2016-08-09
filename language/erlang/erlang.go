package erlang

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])

	// Compile all files except the first
	for _, file := range files[1:] {
		stdout, stderr, err := cmd.Run(workDir, "erlc", file)
		if err != nil {
			return stdout, stderr, err
		}
	}

	// Run first file with escript
	return cmd.RunStdin(workDir, stdin, "escript", files[0])
}
