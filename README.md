## Go AutoF5

AutoF5 is a simple no dependency middleware to automatically reload your browser when go application restarts. It is useful when you are developing a web application and want to see the changes immediately after you rebuild and restart.

It injects a script tag into the html response to reload the browser every time the go application restarts. It does not require any browser extension or any other client-side code.

AutoF5 does not watch the file system or rebuild the go application. You must use other tools similar to [Air](https://github.com/cosmtrek/air), [wgo](https://github.com/bokwoon95/wgo) to rebuild and restart the go application when a file changes.

## Installation

```bash
go get -u github.com/mua/go-autof5
```

## Usage

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/mua/go-autof5"
)

func main() {    
   http.HandleFunc("/", autof5.Livereload(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<html><body><h1>AutoF5</h1></body></html>")
    }))
    http.ListenAndServe(":8080", nil)
}
```