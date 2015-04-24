package c

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.out"

    sourceFiles := filterSourceFiles(files)
    args := append([]string{"clang", "-o", binName}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, binPath)
}

func filterSourceFiles(files []string) []string {
    var newFiles []string

    for _, file := range files {
        if filepath.Ext(file) == ".c" {
            newFiles = append(newFiles, file)
        }
    }

    return newFiles
}
