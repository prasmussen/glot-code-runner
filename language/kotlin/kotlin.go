package kotlin

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    fname := filepath.Base(files[0])

    stdout, stderr, err := cmd.Run(workDir, "kotlinc", fname)
    if err != nil {
        return stdout, stderr, err
    }

    return cmd.Run(workDir, "kotlin", className(fname))
}

func className(fname string) string {
    ext := filepath.Ext(fname)
    return fname[0:len(fname) - len(ext)]
}
