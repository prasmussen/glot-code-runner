package clojure

import (
    "path/filepath"
    "../../cmd"
)

func Run(files []string, stdin string) (string, string, error) {
    workDir := filepath.Dir(files[0])
    return cmd.Run(workDir, "java", "-cp", "/usr/share/java/clojure.jar", "clojure.main", files[0])
}
