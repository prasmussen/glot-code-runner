package language

import (
	"github.com/prasmussen/glot-code-runner/language/assembly"
	"github.com/prasmussen/glot-code-runner/language/ats"
	"github.com/prasmussen/glot-code-runner/language/bash"
	"github.com/prasmussen/glot-code-runner/language/c"
	"github.com/prasmussen/glot-code-runner/language/clojure"
	"github.com/prasmussen/glot-code-runner/language/cobol"
	"github.com/prasmussen/glot-code-runner/language/coffeescript"
	"github.com/prasmussen/glot-code-runner/language/cpp"
	"github.com/prasmussen/glot-code-runner/language/crystal"
	"github.com/prasmussen/glot-code-runner/language/csharp"
	"github.com/prasmussen/glot-code-runner/language/d"
	"github.com/prasmussen/glot-code-runner/language/elixir"
	"github.com/prasmussen/glot-code-runner/language/elm"
	"github.com/prasmussen/glot-code-runner/language/erlang"
	"github.com/prasmussen/glot-code-runner/language/fsharp"
	"github.com/prasmussen/glot-code-runner/language/golang"
	"github.com/prasmussen/glot-code-runner/language/groovy"
	"github.com/prasmussen/glot-code-runner/language/haskell"
	"github.com/prasmussen/glot-code-runner/language/idris"
	"github.com/prasmussen/glot-code-runner/language/java"
	"github.com/prasmussen/glot-code-runner/language/javascript"
	"github.com/prasmussen/glot-code-runner/language/julia"
	"github.com/prasmussen/glot-code-runner/language/kotlin"
	"github.com/prasmussen/glot-code-runner/language/lua"
	"github.com/prasmussen/glot-code-runner/language/mercury"
	"github.com/prasmussen/glot-code-runner/language/nim"
	"github.com/prasmussen/glot-code-runner/language/ocaml"
	"github.com/prasmussen/glot-code-runner/language/perl"
	"github.com/prasmussen/glot-code-runner/language/perl6"
	"github.com/prasmussen/glot-code-runner/language/php"
	"github.com/prasmussen/glot-code-runner/language/python"
	"github.com/prasmussen/glot-code-runner/language/ruby"
	"github.com/prasmussen/glot-code-runner/language/rust"
	"github.com/prasmussen/glot-code-runner/language/scala"
	"github.com/prasmussen/glot-code-runner/language/swift"
	"github.com/prasmussen/glot-code-runner/language/typescript"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"ats":          ats.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"cobol":        cobol.Run,
	"coffeescript": coffeescript.Run,
	"crystal":      crystal.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"elm":          elm.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"groovy":       groovy.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"kotlin":       kotlin.Run,
	"lua":          lua.Run,
	"mercury":      mercury.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"perl6":        perl6.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
	"typescript":   typescript.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
