package scala

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])

    args := append([]string{"scalac"}, files...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    return cmd.Run(workDir, "scala", "Main")
}
