package fsharp

import (
    "path/filepath"
    "../../util"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.exe"

    sourceFiles := util.FilterByExtension(files, "fs")
    args := append([]string{"fsharpc", "--out:" + binName}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, "mono", binPath)
}
