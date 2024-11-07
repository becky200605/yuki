# TASK
## 题目
2. 使用for循环生成20个goroutine，并向一个channel传入随机数和goroutine编号，等待这些goroutine都生成完后，想办法给这些goroutine按照编号进行排序(输出排序前和排序后的结果,要求不使用额外的空间存储着20个数据)  
代码  
```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	type shu struct {
		value int
		order int
	}
	m := make(chan shu, 20)
	//用for循环写出20个goroutine并写入数据
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(order int) {
			m <- shu{rand.Intn(100), i}
			defer wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(m)
	}()
	//读取数据
	var list []shu
	var a shu
	for j := 0; j < 20; j++ {
		a = <-m
		list = append(list, a)
	}
	//输出初始排序
	fmt.Println("初始排序:", list)
	//排序操作
	sort.Slice(list, func(i, j int) bool { return list[i].order < list[j].order })
	//输出操作后排序
	fmt.Println("操作后排序", list)
}
```
输出结果  
```go
初始排序: [{22 0} {98 1} {51 4} {76 2} {11 3} {42 6} {33 5} {62 9} {77 7} {54 8} {36 12} {81 10} {93 11} {89 14} {77 13} {70 17} {84 15} {92 18} {96 16} {51 19}]
操作后排序 [{22 0} {98 1} {76 2} {11 3} {51 4} {33 5} {42 6} {77 7} {54 8} {62 9} {81 10} {93 11} {36 12} {77 13} {89 14} {84 15} {96 16} {70 17} {92 18} {51 19}]
```
