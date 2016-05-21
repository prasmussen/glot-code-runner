package haskell

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    return cmd.Run(workDir, "runghc", files[0])
}
