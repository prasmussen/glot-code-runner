glot-code-runner
================


## Overview
glot-code-runner is a command line application that reads code as a
json payload from stdin – compiles and runs the code – and writes
the result as json to stdout.

## Prerequisites
glot-code-runner requires that the compiler / interpreter for the languages
you want to run is installed and is in PATH.

### Downloads
- [glot-code-runner-linux-x64.tar.gz v1.0.0](https://drive.google.com/uc?id=0B3X9GlR6EmbnOWNITnpZSk9SdVE)

## Supported languages
- assembly
- bash
- c
- clojure
- csharp
- elixir
- cpp
- erlang
- fsharp
- haskell
- golang
- java
- javascript
- lua
- nim
- ocaml
- perl
- php
- python
- ruby
- rust
- scala

## Input (stdin)
The input is required to be a json object containing the properties `language`
and `files`. `language` must be a lowecase string matching one of the supported
languages. `files` must be an array with at least one object containing the
properties `name` and `content`. `name` is the name of the file and can include
forward slashes to create the file in a subdirectory relative to the base
directory. All files are written into the same base directory under the OS's
temp dir.

## Output (stdout)
The output is a json object containing the properties `stdout`, `stderr` and
`error`. `stdout` and `stderr` is captured from the output of the ran code.
`error` is popuplated if there is a compiler / interpreter error.
Note that glot-code-runner will exit with a non-zero code if invalid input is
given or if the files cannot be written to disk (permissions, disk space, etc).
No json will be written to stdout in those cases. Otherwise the exit code is 0.

## Examples
#### Input
    {
      "language": "python",
      "files": [
        {
          "name": "main.py",
          "content": "print(42)"
        }
      ]
    }

#### Output
    {
      "stdout": "42\n",
      "stderr": "",
      "error": ""
    }
