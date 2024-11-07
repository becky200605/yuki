# task
## 并发中出现panic的可能情况
**原代码** 
```go
package main

type message struct {
    Topic     string
        Partition int32
            Offset    int64
            }

            type FeedEventDM struct {
                Type    string
                    UserID  int
                        Title   string
                            Content string
                            }

                            type MSG struct {
                                ms        message
                                    feedEvent FeedEventDM
                                    }

                                    const ConsumeNum = 5

                                    func main() {
                                        var consumeMSG []MSG
                                            var lastConsumeTime time.Time // 记录上次消费的时间
                                                msgs := make(chan MSG)

                                                    //这里源源不断的生产信息
                                                        go func() {
                                                               for i := 0; ; i++ {
                                                                         msgs <- MSG{
                                                                                      ms: message{
                                                                                                      Topic:     "消费主题",
                                                                                                                      Partition: 0,
                                                                                                                                      Offset:    0,
                                                                                                                                                   },
                                                                                                                                                                feedEvent: FeedEventDM{
                                                                                                                                                                                Type:    "grade",
                                                                                                                                                                                                UserID:  i,
                                                                                                                                                                                                                Title:   "成绩提醒",
                                                                                                                                                                                                                                Content: "您的成绩是xxx",
                                                                                                                                                                                                                                             },
                                                                                                                                                                                                                                                       }
                                                                                                                                                                                                                                                                 //每次发送信息会停止0.01秒以模拟真实的场景
                                                                                                                                                                                                                                                                           time.Sleep(100 * time.Millisecond)
                                                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                                                                      }()

                                                                                                                                                                                                                                                                                          //不断接受消息进行消费
                                                                                                                                                                                                                                                                                              for msg := range msgs {
                                                                                                                                                                                                                                                                                                     // 添加新的值到events中
                                                                                                                                                                                                                                                                                                            consumeMSG = append(consumeMSG, msg)
                                                                                                                                                                                                                                                                                                                   // 如果数量达到额定值就批量消费
                                                                                                                                                                                                                                                                                                                          if len(consumeMSG) >= ConsumeNum {
                                                                                                                                                                                                                                                                                                                                    //进行异步消费
                                                                                                                                                                                                                                                                                                                                              go func() {
                                                                                                                                                                                                                                                                                                                                                           m := consumeMSG[:ConsumeNum]
                                                                                                                                                                                                                                                                                                                                                                        fn(m)
                                                                                                                                                                                                                                                                                                                                                                                  }()
                                                                                                                                                                                                                                                                                                                                                                                            // 更新上次消费时间
                                                                                                                                                                                                                                                                                                                                                                                                      lastConsumeTime = time.Now()
                                                                                                                                                                                                                                                                                                                                                                                                                // 清除插入的数据
                                                                                                                                                                                                                                                                                                                                                                                                                          consumeMSG = consumeMSG[ConsumeNum:]
                                                                                                                                                                                                                                                                                                                                                                                                                                 } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
                                                                                                                                                                                                                                                                                                                                                                                                                                           // 如果距离上次消费已经超过5分钟且有未处理的消息
                                                                                                                                                                                                                                                                                                                                                                                                                                                     if len(consumeMSG) > 0 {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  //进行异步消费 
                                                                                                                                                                                                                                                                                                                                                                                                                                                                               go func() {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               m := consumeMSG[:ConsumeNum]
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               fn(m)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            }()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         // 更新上次消费时间
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      lastConsumeTime = time.Now()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   // 清空插入的数据
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                consumeMSG = consumeMSG[ConsumeNum:]
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 }
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     }

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     func fn(m []MSG) {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         fmt.Printf("本次消费了%d条消息\n", len(m))
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         }
     
```
出现panic的报错
```go
panic: runtime error: slice bounds out of range [:5] with capacity 3

goroutine 20 [running]:
main.main.func2()
        D:/go/panic.go:61 +0x3c
created by main.main in goroutine 1
        D:/go/panic.go:60 +0x3db
exit status 2
```
由报错可知，出现panic的原因是对切片的访问超出索引范围，即出现在
```go
// 如果数量达到额定值就批量消费
       if len(consumeMSG) >= ConsumeNum {
          //进行异步消费
          go func() {
             m := consumeMSG[:ConsumeNum]
             fn(m)
          }()
          // 更新上次消费时间
          lastConsumeTime = time.Now()
          // 清除插入的数据
          consumeMSG = consumeMSG[ConsumeNum:]
          } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
          // 如果距离上次消费已经超过5分钟且有未处理的消息
          if len(consumeMSG) > 0 {
             //进行异步消费
             go func() {
                m := consumeMSG[:ConsumeNum]
                fn(m)
             }()
             // 更新上次消费时间
             lastConsumeTime = time.Now()
             // 清空插入的数据
             consumeMSG = consumeMSG[ConsumeNum:]
          }
       }
```
出现越界的原因:切片的并发修改导致切片长度存在比ConsumeMSG小的可能  
**解决方案**  
利用互斥锁保证对切片的操作可以正常进行  
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)
	var mu sync.Mutex
	//这里源源不断的生产信息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}
			//每次发送信息会停止0.01秒以模拟真实的场景
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//不断接受消息进行消费
	for msg := range msgs {
      //加锁，保证其他协程不改变consumeMSG
		mu.Lock()
		// 添加新的值到events中
		consumeMSG = append(consumeMSG, msg)
		// 如果数量达到额定值就批量消费

		if len(consumeMSG) >= ConsumeNum {
			//进行异步消费
			go func() {

				m := consumeMSG[:ConsumeNum]
				fn(m)

				// 更新上次消费时间
				lastConsumeTime = time.Now()
            //加锁保证数据清楚正常进行
				mu.Lock()
				// 清除插入的数据
				consumeMSG = consumeMSG[ConsumeNum:]
				mu.Unlock()
			}()
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			// 如果距离上次消费已经超过5分钟且有未处理的消息
			if len(consumeMSG) > 0 {
				//进行异步消费
				go func() {

					var m []MSG
               //保证切片不越界
					if len(consumeMSG) >= ConsumeNum {
						m = consumeMSG[:ConsumeNum] // 这里会从 consumeMSG 中获取前 ConsumeNum 条消息
					} else {
						m = consumeMSG[:]
					}
					fn(m)

					// 更新上次消费时间
					lastConsumeTime = time.Now()
					//加锁保证数据清楚正常进行
					mu.Lock()
               // 清除插入的数据
					consumeMSG = consumeMSG[ConsumeNum:] // 这里把 consumeMSG 的前 ConsumeNum 条消息去除
					mu.Unlock()
				}()

				// 更新上次消费时间
				lastConsumeTime = time.Now()
				// 清空插入的数据

				consumeMSG = consumeMSG[ConsumeNum:]

			}
		}
		mu.Unlock()
	}
}

func fn(m []MSG) {
   fmt.Printf("本次消费了%d条消息\n", len(m))
}
```

