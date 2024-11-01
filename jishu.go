package main
import "fmt"
func jishu() func(x int) int{
	return func (x int) int{
		return x+1
	}
}
func main(){
	var a,b int
	count :=jishu()
	fmt.Scanf("%d",&b)
	//计数器：如果输入b是1，则计数器+1
	if(b==1){
		a=count()
	}
	fmt.Printf(a)
}
