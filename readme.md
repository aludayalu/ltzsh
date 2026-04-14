# LTZ

LTZ is a work environment for Lumatozer.

## Compiling

Aggressive Inlining and Bounds Checking Disabled

```go
go build -gcflags="all=-l=4 -B"
```


### With binary size optimizations.

**Note:** The `-s` and `-w` ldflags still do not remove the stack track information on crashes or errors.
```go
go build -trimpath -gcflags="all=-l=4 -B" -ldflags="-s -w"
```