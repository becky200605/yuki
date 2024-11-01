# task2
** 比较字符串"hello，世界"的长度 和for range 该字符串的循环次数 **
由题意可设计如下代码
```go
package main
import "fmt"
func main(){
  var i string
  var a,b int
  i="hello,世界"
  //长度
  a=len(i)
  //for range
  for range i{
  b++
  }
  fmt.Printf("%d,%d",a,b)
}
```
有输出结果
```go
12,8
```

