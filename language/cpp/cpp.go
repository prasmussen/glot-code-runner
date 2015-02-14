package cpp

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    binName := "a.out"

    stdout, stderr, err := cmd.Run(workDir, "clang++", files[0], "-o", binName)
    if err != nil {
        return stdout, stderr, err
    }

    binPath := filepath.Join(workDir, binName)
    return cmd.Run(workDir, binPath)
}
