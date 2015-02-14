package cmd

import (
    "bytes"
    "os/exec"
)

func Run(workDir string, args ...string) (string, string, error) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer

    cmd := exec.Command(args[0], args[1:]...)
    cmd.Dir = workDir
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()

    return stdout.String(), stderr.String(), err
}
