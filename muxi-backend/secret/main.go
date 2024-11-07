package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"muxi-backend/tool/getDecryptedPaper"
	"muxi-backend/tool/savePaper"
	"net/http"
)

func main() {
	// 目标根URL
	url := "http://121.43.151.190:8000/"
	// 发送 GET 请求,返回的结果还需要进行处理才能得到你需要的结果
	response1, err := http.Get(url + "paper")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response1.Body.Close()
	EncodedPaper, err := ioutil.ReadAll(response1.Body)
	if err != nil {
		fmt.Println("Read from response1.Body Failed,err:", err)
		return
	}
	response2, err := http.Get(url + "secret")
	if err != nil {
		log.Fatalf("Failed to send request :%v", err)
	}
	defer response2.Body.Close()
	Key, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		fmt.Println("Read from response2.Body Failed,err:", err)
		return
	}
	Text := getDecryptedPaper.GetDecryptedPaper(string(EncodedPaper), string(Key))
	Path := "D:/muxi-backend/paper/Academician Sun's papers"
	savePaper.SavePaper(Path, Text)

}
