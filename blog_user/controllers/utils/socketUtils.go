package utils

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 1024,
		// 写入存储空间大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var (
	wbsCon *websocket.Conn
	err    error
	data   []byte
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 完成http应答，在httpheader中放下如下参数
	if wbsCon, err = upgrader.Upgrade(w, r, nil); err != nil {
		return // 获取连接失败直接返回
	}
	go receiveMessage()
}

func StartSocket() {
	// 当有请求访问ws时，执行此回调方法
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}

func CloseSocket() {
	wbsCon.Close()
}

func receiveMessage() {
	for {
		// 只能发送Text, Binary 类型的数据,下划线意思是忽略这个变量.
		if _, data, err = wbsCon.ReadMessage(); err != nil {
			CloseSocket()
		}
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			CloseSocket()
		}
	}
}

func SendSocketMessage(message string)  {
	if err = wbsCon.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		CloseSocket()
	}
}
