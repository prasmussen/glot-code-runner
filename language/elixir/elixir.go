package elixir

import (
    "path/filepath"
    "../../cmd"
    "../../util"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    sourceFiles := util.FilterByExtension(files, "ex")
    args := append([]string{"elixirc"}, sourceFiles...)
    return cmd.RunStdin(workDir, stdin, args...)
}
