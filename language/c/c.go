package c

import (
    "path/filepath"
    "../../util"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.out"

    sourceFiles := util.FilterByExtension(files, "c")
    args := append([]string{"clang", "-o", binName, "-lm"}, sourceFiles...)
    stdout, stderr, err := cmd.Run(workDir, args...)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, binPath)
}
