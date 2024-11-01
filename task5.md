# task5
**以下哪里x进行重新声明，为什么**
```go

  func main() {
      x := "hello!"
      for i := 0; i < len(x); i++ {
          x := x[i]
          if x != '!' {
              x := x + 'A' - 'a'
              fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
         }
     }
  }
```
A：x在for循环内，if语句内都有重新声明，因为出现了**：=**的声明，每次重新声明的x作用域都不同。

