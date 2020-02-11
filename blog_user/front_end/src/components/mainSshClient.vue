<template>
  <div id="content">
    <div class="fontsize">
      <button class="fontsize" @click="sshbuild">Build</button>
</div>
</div>
</template>


<script>
    export default {
        methods: {
            sshbuild: function () {
                this.$axios.post(global.sshbuild,{
                    'articleId' : 123
                }).then(
                    response => {
                        this.article = JSON.parse(response.data)["data"]
                    },
                    response => console.log(response)
                )
            }
        }
    }
</script>


<style>
  .fontsize {
    display: inline-block;
    color: #222;
    font-size: 26px;
    font-weight: 600;
    border-bottom: 1px solid white;
    cursor: pointer;
  }
</style>

