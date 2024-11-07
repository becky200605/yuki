# task
**题目**  
3. 经典老题：交叉打印下面两个字符串（要求一个打印完，另一个会继续打印）
"ABCDEFGHIJKLMNOPQRSTUVWXYZ" "0123..."
得到："AB01CD23EF34..."
代码如下：
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var char []rune
	var nums []int
	var k, q int
	var wg sync.WaitGroup
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)
	for i := 'A'; i <= 'Z'; i++ {
		char = append(char, i)
		k++
	}
	for j := 0; j < 26; j++ {
		nums = append(nums, j)
	}
	wg.Add(2)
	go func() {

		defer wg.Done()

		for a := 0; a < len(char); a += 2 {
			mu.Lock()
			//用q的奇偶性来控制下一个打的是字母还是数字
			if q%2 != 0 {
				cond.Wait()
			}
			fmt.Printf("%c%c", char[a], char[a+1])
			q++
			//唤醒被阻塞的数字协程
			cond.Signal()
			mu.Unlock()

		}
	}()

	go func() {

		defer wg.Done()
		for b := 0; b < len(nums); b += 2 {
			mu.Lock()
			if q%2 == 0 {
				cond.Wait()
			}
			fmt.Printf("%d%d", nums[b], nums[b+1])
			q++
			//唤醒被阻塞的字母协程
			cond.Signal()
			mu.Unlock()

		}
	}()
	wg.Wait()
}
```
输出如下
```go
AB01CD23EF45GH67IJ89KL1011MN1213OP1415QR1617ST1819UV2021WX2223YZ2425
```
