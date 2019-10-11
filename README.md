[![GoDoc](https://godoc.org/github.com/nickwells/twrap.mod?status.png)](https://godoc.org/github.com/nickwells/twrap.mod)

# twrap.mod
This provides ways of wrapping and indenting text printed to the terminal.

You first construct a wrapper object TWConf and then you can use the various
Wrap methods to print text indented and wrapped at the target line length.

Here is a simple example

```go
twc, _ := twrap.NewTWConf()
twc.Wrap("the quality of mercy is not strained", 5)
```

More examples are available in the example test files.

There are also various `Print` functions which simply call `fmt.Fprint` functions and allow you to write:

```go
twc.Print("Hello")
```

rather than:

```go
fmt.Fprint(twc.W, "Hello")
```
