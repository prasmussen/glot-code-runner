package language

import (
    "./assembly"
    "./bash"
    "./c"
    "./clojure"
    "./coffeescript"
    "./csharp"
    "./d"
    "./elixir"
    "./cpp"
    "./erlang"
    "./fsharp"
    "./haskell"
    "./idris"
    "./golang"
    "./java"
    "./javascript"
    "./julia"
    "./lua"
    "./nim"
    "./ocaml"
    "./perl"
    "./php"
    "./python"
    "./ruby"
    "./rust"
    "./scala"
    "./swift"
)

type runFn func([]string) (string, string, error)

var languages = map[string]runFn{
    "assembly": assembly.Run,
    "bash": bash.Run,
    "c": c.Run,
    "clojure": clojure.Run,
    "coffeescript": coffeescript.Run,
    "csharp": csharp.Run,
    "d": d.Run,
    "elixir": elixir.Run,
    "cpp": cpp.Run,
    "erlang": erlang.Run,
    "fsharp": fsharp.Run,
    "haskell": haskell.Run,
    "idris": idris.Run,
    "go": golang.Run,
    "java": java.Run,
    "javascript": javascript.Run,
    "julia": julia.Run,
    "lua": lua.Run,
    "nim": nim.Run,
    "ocaml": ocaml.Run,
    "perl": perl.Run,
    "php": php.Run,
    "python": python.Run,
    "ruby": ruby.Run,
    "rust": rust.Run,
    "scala": scala.Run,
    "swift": swift.Run,
}

func IsSupported(lang string) bool {
    _, supported := languages[lang]
    return supported
}

func Run(lang string, files []string) (string, string, error) {
    return languages[lang](files)
}
