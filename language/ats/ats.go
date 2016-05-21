package ats

import (
    "path/filepath"
    "../../util"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.out"

    sourceFiles := util.FilterByExtension(files, "dats")
    args := append([]string{"patscc", "-o", binName}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.RunStdin(workDir, stdin, binPath)
}
