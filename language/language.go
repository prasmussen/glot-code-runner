package language

import (
    "./assembly"
    "./bash"
    "./c"
    "./clojure"
    "./csharp"
    "./elixir"
    "./cpp"
    "./erlang"
    "./fsharp"
    "./haskell"
    "./golang"
    "./java"
    "./javascript"
    "./lua"
    "./ocaml"
    "./perl"
    "./php"
    "./python"
    "./ruby"
    "./rust"
    "./scala"
)

type runFn func([]string) (string, string, error)

var languages = map[string]runFn{
    "assembly": assembly.Run,
    "bash": bash.Run,
    "c": c.Run,
    "clojure": clojure.Run,
    "csharp": csharp.Run,
    "elixir": elixir.Run,
    "cpp": cpp.Run,
    "erlang": erlang.Run,
    "fsharp": fsharp.Run,
    "haskell": haskell.Run,
    "go": golang.Run,
    "java": java.Run,
    "javascript": javascript.Run,
    "lua": lua.Run,
    "ocaml": ocaml.Run,
    "perl": perl.Run,
    "php": php.Run,
    "python": python.Run,
    "ruby": ruby.Run,
    "rust": rust.Run,
    "scala": scala.Run,
}

func IsSupported(lang string) bool {
    _, supported := languages[lang]
    return supported
}

func Run(lang string, files []string) (string, string, error) {
    return languages[lang](files)
}
