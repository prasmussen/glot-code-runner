package ocaml

import (
    "path/filepath"
    "../../util"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a"

    sourceFiles := reverse(util.FilterByExtension(files, "ml"))
    args := append([]string{"ocamlc", "-o", binName}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, binPath)
}

func reverse(files []string) []string {
    reversed := make([]string, 0, len(files))

    for i := len(files) - 1; i >= 0; i-- {
        reversed = append(reversed, files[i])
    }

    return reversed
}
