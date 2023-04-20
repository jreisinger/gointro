Go intro for sec and ops people
-------------------------------

Why intro for sec and ops people

* since 2013 I've worked as sysadmin, devops/security person
* since 2008 I've been using Go (together with some Bash and Perl/Python)
* mostly small/medium sized one/two-person programs

Why sec and ops people should program

* to better understand how software works
* to better understand dev people
* to create tools, automation, PoCs

Why Go

* focus on simplicity, reliability (correctness) and speed
* lot of cloud is written in Go (e.g. Kubernetes, terraform)
* nice addition to Bash and/or Python since it's quite different

Go peculiarities

* packages, types, goroutines, channels: `cd fetch`
* cross-compilation to a single binary: `GOOS=linux GOARCH=arm64 go build fetch.go`
* methods and interfaces: `cd shop`
* small and secure images: `docker build -t shop . && docker run -p 8000:8000 shop`

Learn more

* books: Learning Go, The Go Programming Language, Black Hat Go
* online: https://go.dev/tour, https://gobyexample.com/, https://go.dev/doc
* https://github.com/jreisinger/gokatas