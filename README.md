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

## 类型断言

### x.(T) 

x是接口类型的表达式，T是一个类型（可以是一个具体类型，也可以是另一个接口）

### x.(type)

```go
switch x.(type) {
    case nil: //
    case int,uint: //
    case bool: //
    case string: //
    default: //
}
```

## 反射

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

## 共享变量实现并发

golang的准则：在并发环境下，不要通过共享内存来通信，而要通过通道。

### 互斥锁
```go

// 互斥锁
sync.Mutex

// 通过使用
defer mu.Unlock()
// 来确保解锁

// go语言的互斥量是不可重入的。

// 读写互斥锁
sync.RWMutex

// 延迟初始化
sync.Once

```
### 竞态检测器

## 方法

- 可以给任意类型定义方法，并不一定是结构体：
```go
type Path []Point
func (path Path) Distance() float64 {
    ...
}
```

- 方法的指针接收者，可以避免复制实参
- 可以调用内嵌结构体的方法，其实这是编译器实现的桥接，并不是继承。
- 定义一个“方法”变量
```go
var op func(p,q Point) Point
```

## 封装

- go语言只有一种封装，就是名字的大小写，首字母大写的可以从包中导出。
- 要封装一个对象，必须使用结构体
- 封装的单元是包而不是类型，无论是函数还是方法，结构体类型里的字段对于同一个包里的所有代码都是可见的。

## 调试 debug (delve)


```delve

config -list

config max-string-len 1000

config max-array-values 1000

```

## 覆盖率测试 test coverage

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go test -covermode=count -coverprofile=count.out fmt
go tool cover -html=count.out

```

## module

### 找到引入 indirect 依赖的模块

```shell
go mod why -m $indirect-module-name
```
