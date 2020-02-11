<template>
  <div id="content" class="fontsize">
    <button class="fontsize" @click="sshbuild">Build</button>
    <button class="fontsize" @click="stopBuild">停止编译</button>
    <button class="fontsize" @click="connectSocket">连接</button>
    <button class="fontsize" @click="sendSocketMsg">发消息</button>
    <div class="fontsize" id="buildProgress">
      <h1>编译进度 : {{buidProgress}}</h1>
    </div>
  </div>
</template>


<script>
    let ws;
    export default {
        data() {
            return {
                buidProgress: "未开始",
            }
        },
        mounted: function () {

            ws = new WebSocket("ws://127.0.0.1:8888/ws");//连接服务器
            ws.onopen = function (event) {
                console.log(event);
                alert('连接了');
            };
            ws.onmessage = function (event) {
                let date = new Date();
                let msg = date.toLocaleString() + event.data;
                if (msg.indexOf("buildLog") >= 0) {
                    set(this.buidProgress, 'name');
                } else {
                    alert(msg);
                }
            }
            ws.onclose = function (event) {
                alert("已经与服务器断开连接\r\n当前连接状态：" + this.readyState);
            };
            ws.onerror = function (event) {
                alert("WebSocket异常！");
            };
        },
        methods: {
            sshbuild: function () {
                this.$axios.post(global.sshbuild, {
                    'articleId': 123
                }).then(
                    response => {
                        this.article = JSON.parse(response.data)["data"]
                    },
                    response => console.log(response)
                )
            },
            connectSocket: function () {
                ws = new WebSocket("ws://127.0.0.1:8888/ws");//连接服务器
                ws.onopen = function (event) {
                    console.log(event);
                    alert('连接了');
                };
                ws.onmessage = function (event) {
                    let date = new Date();
                    let msg = date.toLocaleString() + event.data;
                    if (msg.indexOf("buildLog") >= 0) {
                        this.buidProgress = msg
                    } else {
                        alert(msg);
                    }
                }
                ws.onclose = function (event) {
                    alert("已经与服务器断开连接\r\n当前连接状态：" + this.readyState);
                };
                ws.onerror = function (event) {
                    alert("WebSocket异常！");
                };
            },
            sendSocketMsg: function () {
                ws.send("hello");
            },
            stopBuild: function () {

            },
        }
    }
</script>


<style>
  .fontsize {
    flex-direction:column;
    display: inline-block;
    color: #222;
    font-size: 26px;
    font-weight: 600;
    border-bottom: 1px solid white;
    cursor: pointer;
  }
</style>

