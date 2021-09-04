# go-experiment

## basic 
- 每一个文件夹定义一个package
- 名为main的pakcage表示它包含了main函数，用于执行，而不是一个包
- main function是程序的入口
```go
package main

import (
    "fmt"
    "os"
)

func main() {

}
```

## 循环

- for 是唯一的循环语句
- 每一次迭代，range产生一对值

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

## const
- 圆括号表示并列的关系
```go
const (
    white = 0
    black = 1
)
```

## 初始化
```go
s := ""
var s string
var s = ""
var s string = ""
```
## slice
```go
var palette = []color.Color{color.White, color.Black}
```

## map
```go

counts := make(map[string]int)
// counts 是一个引用，指向map
```

## log
```go
log.Fatal("")
```

### 类型断言

#### x.(T) 

x是接口类型的表达式，T是一个类型（可以是一个具体类型，也可以是另一个接口）

#### x.(type)

```go
switch x.(type) {
    case nil: //
    case int,uint: //
    case bool: //
    case string: //
    default: //
}
```

#### 反射

```go
reflect.TypeOf(i interface{})
reflect.ValueOf(i interface{})

func Any(value interface{}) string {
    return formatAtom(refect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
    switch v.Kind() {
    case reflect.Int64:
    case reflect.Bool:
    case reflect.Invalid:
    default:
        return v.Type().String() + " value"
    }
}

```

传入任意类型，被隐式转换为空接口 interface{}, 因为接口具有两个成员：具体类型，具体值。

反射最常见的应用场景：
数据在文件和内存之间的序列号和反序列化。

参考: https://www.cnblogs.com/qcrao-2018/p/10822655.html