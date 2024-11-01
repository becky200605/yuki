# task4
**如何比较结构体**  
1.直接比较  
一般情况下，结构体可以利用关系运算符**==**或**！=**进行比较  
2.手动比较  
对于一些不可直接比较的，可以定义一个比较函数进行比较  
eg:
```go
package main
import "fmt"
func main(){
type jiegou struct{
  a int
  b []int
  }
func main(){
  i:=jiegou{a:1,b:[]int{1,2,3}}
  j:=jirgou{a:2,b:[]int{1,2}}
  if(i.a<j.a||len(i.b)<len(j.b)){
  fmt.Println("a<b")
 }else {
   fmt.Println("a>b")
}
```
3.使用函数  
更复杂的情况下，可以利用<table><tr><tb bgcolor=blue>reflect.DeepEqual</td></tr></table>函数进行结构体的比较  
## reflect.DeepEqual
对于一些切片、字典、数组和结构体等类型，想要比较两个值是否相等，可以用这个函数  
**原理**  
1.首先检查传入的两个参数是否为nil，若有则直接返回false  
2.如果两个参数是相同类型的值类型或引用类型，则直接比较二进制中的值  
3.如果类型不同，则先将其转换为interface()类型，再比较  



