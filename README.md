# go-experiment

- basic 
```go
package main

import (
    "fmt"
    "os"
)

func main() {

}
```

- for 是唯一的循环语句

```go
for init; condition; post {

}

// "while"
for condition {

}

// endless loop
for {
    
}

for index, v := range a_slice {

}

for k,v := range a_map {

}
```
- const
```go
const (
    white = 0
    black = 1
)
```

- init
```go
s := ""
var s string
var s = ""
var s string = ""
```
- slice
```go
var palette = []color.Color{color.White, color.Black}
```
- map
```go

counts := make(map[string]int)
// counts 是一个引用，指向map
```

- log
```go
log.Fatal("")
```
