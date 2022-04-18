action
```go
```

create a new module
```go
go mod init github.com/org/my-module
```
run go program
```go
go run .
```
reload module
```go
// The -u flag updates modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.
go get -u
go mod tidy
```
helloworld app
```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello", "World!")
	fmt.Println(quote.Go())
}
```