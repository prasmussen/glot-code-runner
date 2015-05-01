package lua

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    return cmd.Run(workDir, "lua", files[0])
}
