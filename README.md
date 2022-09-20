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

## package

一个 package 就是一个文件夹，里面必须有一个 .go 文件，可以导出给其他地方使用

一个或者多个 package（文件夹） 可以组成一个 module, module 的根目录下必须包含 go.mod 和 go.sum，module 是为了解决 package 的版本以及依赖的。之前的解决方案是 GOPATH 和 vendor

多个 module 可以在同一个 repository 里面

## 封装

- go语言只有一种封装，就是名字的大小写，首字母大写的可以从包中导出。
- 要封装一个对象，必须使用结构体
- 封装的单元是包而不是类型，无论是函数还是方法，结构体类型里的字段对于同一个包里的所有代码都是可见的。


## Module

### Graph

```bash
go mod graph 
```

的结果，不区分直接依赖和间接依赖。即：

main module --> dependency module 这里包括所有的直接和间接依赖。

### Get

从 go1.18 起，get 只更新 go.mod 和 go.sum 这两个文件，不会再做 build，也不会去安装 binary（go install 用来安装）

```bash
go get ${package_name}
```
更新 package 到最新的满足兼容要求的版本

```bash
go get -u ${package_name}
```
更新package 所在的 module 的版本，但是并不做编译检查，要使用 go mod tidy 做清理和编译检查。

### Tidy

```bash
go mod tidy
```
会检查 package 是否在 module 里，并不升级依赖的版本

### Why

```bash
go mod why -m ${module_name}
```
检查 module 所在的 package 的最短依赖

### List

```bash
go list -m -u ${module_name}
```
查看当前生效的版本以及可以升级的版本


### 其他

老的项目，没有使用 module 的话，上面的命令都看不到有用的结果。需要的话，到 vendor 目录去看。

## cgo 和 goc

- [command cgo](https://golang.org/cmd/cgo/)
- [C?Go?Cgo!](https://blog.golang.org/c-go-cgo)

## go build 做了什么

```bash
go build -work -a -x -p 1

-work: go build creates a temporary folder for work files. This arg will print out the location of that folder and not delete it after the build
-a: Golang caches previously built packages. -a makes go build ignore the cache so our build will print all steps
-p 1: This sets the concurrency to a single thread to log output linear
-x: go build is a wrapper around other Golang tools like compile. -x outputs the commands and arguments that are sent to these tools
```

根据 log 可以看出，每一个依赖包都会被 compile 成 .a ，然后重命名到一个临时文件，最后全部 link 成一个可执行文件

参考：
- [How go build works](https://maori.geek.nz/how-go-build-works-750bb2ba6d8e)
- [编译成静态库](https://developpaper.com/compilation-and-use-of-go-language-static-library/)

