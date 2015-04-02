package language

import (
    "./bash"
    "./c"
    "./cpp"
    "./erlang"
    "./haskell"
    "./golang"
    "./java"
    "./javascript"
    "./perl"
    "./php"
    "./python"
    "./ruby"
)

type runFn func([]string) (string, string, error)

var languages = map[string]runFn{
    "bash": bash.Run,
    "c": c.Run,
    "cpp": cpp.Run,
    "erlang": erlang.Run,
    "haskell": haskell.Run,
    "go": golang.Run,
    "java": java.Run,
    "javascript": javascript.Run,
    "perl": perl.Run,
    "php": php.Run,
    "python": python.Run,
    "ruby": ruby.Run,
}

func IsSupported(lang string) bool {
    _, supported := languages[lang]
    return supported
}

func Run(lang string, files []string) (string, string, error) {
    return languages[lang](files)
}
