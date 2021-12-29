# jp-dimension-parser
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Test](https://github.com/k-yomo/jp-dimension-parser/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/k-yomo/jp-dimension-parser/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/k-yomo/jp-dimension-parser/branch/main/graph/badge.svg)](https://codecov.io/gh/k-yomo/jp-dimension-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/k-yomo/jp-dimension-parser)](https://goreportcard.com/report/github.com/k-yomo/jp-dimension-parser)


## Installation
```shell
$ go get -u github.com/k-yomo/jp-dimension-parser
```

## Example
```go
import (
    "fmt"
	
    "github.com/k-yomo/jp-dimension-parser/dimparser"
)

func main() {
    dimension := dimparser.Parse("幅62cm×奥行73cm×高さ189cm")
    if dimension != nil {
        fmt.Println("width(cm):", dim.Width.Centimeter())
        fmt.Println("depth(cm):", dim.Depth.Centimeter())
        fmt.Println("height(cm):", dim.Height.Centimeter())
        // => width(cm): 62 
        // => depth(cm): 73 
        // => height(cm): 189
    }
}
```
