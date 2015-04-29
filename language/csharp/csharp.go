package csharp

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.exe"

    sourceFiles := filterSourceFiles(files)
    args := append([]string{"mcs", "-out:" + binName}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, "mono", binPath)
}

func filterSourceFiles(files []string) []string {
    var newFiles []string

    for _, file := range files {
        if filepath.Ext(file) == ".cs" {
            newFiles = append(newFiles, file)
        }
    }

    return newFiles
}
