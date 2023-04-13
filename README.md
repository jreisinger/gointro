Go intro for sec and ops people
-------------------------------

Why sec and ops people should program

* to better understand how software works
* to better understand dev people
* to create tools, automation, PoCs

Why Go

* focus on simplicity, reliability (correctness) and speed
* lot of cloud is written in Go (e.g. Kubernetes)
* nice addition to Bash and Python since it's compiled and peculiar

Go peculiarities

* packages, types, interfaces, pointers, methods (`bytecounter`)
* concurrency via goroutines and channels, defer (`fetch`)
* cross-compilation to a single binary: `GOOS=linux GOARCH=arm64 go build fetch.go`

Learn more

* online: https://go.dev/tour, https://gobyexample.com/, https://go.dev/doc
* books: Learning Go, The Go Programming Language, Black Hat Go
* write tools, automation, PoCs
* https://github.com/jreisinger/gokatas
