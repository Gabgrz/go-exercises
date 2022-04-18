action
```go
```

create a new module
```go
go mod init github.com/my-module
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
share modules
```go
```
