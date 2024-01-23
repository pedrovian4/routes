# Route Package for Go
### Introduction
The route package provides a simple and flexible way to handle routing in Go web applications. It allows for easy registration of routes, application of middlewares, and has utilities for JSON response handling.

### Installation
To install the package, use the following go get command:

```bash
go get github.com/yourusername/route
```

## Usage

```go
package main

import (
	"github.com/pedrovian4/some-app/api/middleware"
	"github.com/pedrovian4/routes"
	"net/http"
)

func main() {
	router := route.NewRouter()

	router.Use(middleware.Log)

	router.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		route.JSON(w, http.StatusOK, map[string]string{"message": "welcome to sports bets"})
	})

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
```


## Middleware



```go
package middleware

import (
	"fmt"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

```