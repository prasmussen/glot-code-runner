package elm

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])

    // Move bootstrap files into work dir
    stdout, stderr, err := cmd.RunBash(workDir, "mv -f /bootstrap/* .")
    if err != nil {
        return stdout, stderr, err
    }

    // Compile elm to javascript
    stdout, stderr, err = cmd.Run(workDir, "elm-make", files[0], "--output", "raw.js")
    if err != nil {
        return stdout, stderr, err
    }

    // Convert to console app
    stdout, stderr, err = cmd.Run(workDir, "bash", "elm-io.sh", "raw.js", "main.js")
    if err != nil {
        return stdout, stderr, err
    }

    // Run javascript with node
    return cmd.RunStdin(workDir, stdin, "node", "main.js")
}
