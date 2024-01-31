[![codecov](https://codecov.io/gh/gofika/httputil/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/httputil)
[![Build Status](https://github.com/gofika/httputil/workflows/build/badge.svg)](https://github.com/gofika/httputil)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/httputil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/httputil)](https://goreportcard.com/report/github.com/gofika/httputil)
[![Licenses](https://img.shields.io/github/license/gofika/httputil)](LICENSE)

# httputil

golang http utils for common use


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/httputil
```

### Example

```go
package main

import (
	"fmt"

	"github.com/gofika/httputil"
)

func main() {
	resp, err := httputil.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
  }
	type GetResp struct {
		URL string `json:"url"`
	}
	res, err := httputil.ReadJSON[GetResp](resp)
	fmt.Printf("url=%s\n", res.URL)
}
```