<template xmlns="http://www.w3.org/1999/html">
  <div id="content" class="fontsize">
    <button class="fontsize" @click="sshbuild">Build</button>
    <br/>
    <button class="fontsize" @click="stopBuild">停止编译</button>
    <br/>
    <button class="fontsize" @click="connectSocket">{{connectStatus}}</button>
    <br/>
    <button class="fontsize" @click="sendSocketMsg">发消息</button>
    <br/>
    <h1>编译进度 : {{progress}}</h1>

    <br/>
    <h1>LOG : {{buildLog}}</h1>
  </div>
</template>


<script>
    const unstartProgress = "未开始";
    const unConnectStatus = "未连接";
    const connectingStatus = "已连接";
    let ws;
    export default {
        data() {
            return {
                progress: unstartProgress,
                connectStatus: unConnectStatus,
                buildLog: "",
            }
        },
        mounted: function () {
            this.connectSocket();
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
                let vm = this;
                if (vm.connectStatus == connectingStatus){
                    return;
                }
                ws = new WebSocket("ws://127.0.0.1:8888/ws");//连接服务器
                ws.onopen = function (event) {
                    vm.connectStatus = connectingStatus
                    // console.log(event);
                    // alert('连接了');
                };
                ws.onmessage = function (event) {
                    let date = new Date();
                    let msg = date.toLocaleString() + event.data;
                    vm.buildLog = msg;
                    console.log("receive " + msg);
                    if (msg.indexOf("buildLog") >= 0) {
                        if (vm.progress == unstartProgress) {
                            vm.progress = "0%";
                        } else if (vm.getProgress(msg).indexOf("%") >= 0) {
                            vm.progress = vm.getProgress(msg);
                        } else if (msg.indexOf("_release END") >= 0) {
                            vm.progress = "100%";
                        }
                    } else {
                        // alert(msg);
                    }
                };
                ws.onclose = function (event) {
                    vm.progress = unstartProgress;
                    vm.connectStatus = unConnectStatus;
                    vm.buildLog = "";
                    // vm.connectSocket();
                    // alert("已经与服务器断开连接\r\n当前连接状态：" + this.readyState);
                };
                ws.onerror = function (event) {
                    vm.progress = unstartProgress;
                    vm.connectStatus = unConnectStatus;
                    vm.buildLog = "";
                    // vm.connectSocket();
                    // alert("WebSocket异常！");
                };
            },
            sendSocketMsg: function () {
                ws.send("hello");
            },
            stopBuild: function () {
                ws.send("stop build");
            },
            getProgress: function (buildLog) {
                let progress = "";
                if (buildLog.indexOf("%") >= 0 && buildLog.indexOf("[") >= 0) {
                    progress = buildLog.split("]")[0].split("[")[1];
                }
                return progress;
            },
        }
    }
</script>


<style>
  .fontsize {
    flex-direction: column;
    display: inline-block;
    color: #222;
    font-size: 26px;
    font-weight: 600;
    border-bottom: 1px solid white;
    cursor: pointer;
  }
</style>

