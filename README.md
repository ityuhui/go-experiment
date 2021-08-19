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

- 每一个文件夹定义一个package
- 名为main的pakcage表示它包含了main函数，用于执行，而不是一个包
- main function是程序的入口
- 圆括号表示并列的关系
- 每一次迭代，range产生一对值
- 创建map:
```
counts := make(map[string] int)
```
