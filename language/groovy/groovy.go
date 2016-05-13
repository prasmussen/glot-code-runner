package groovy

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    return cmd.Run(workDir, "groovy", files[0])
}
