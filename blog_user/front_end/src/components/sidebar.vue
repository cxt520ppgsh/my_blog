<template>
  <div id="left_menu">
    <div class="menu_name">
      <h1>吴亦堂的个人博客</h1>
      <h5>&nbsp;生活本就只有方向没有答案&nbsp;</h5>
    </div>
    <nav>
      <ul>
        <router-link to="/article">
          <li class="nav">
            <span>首页</span>
          </li>
        </router-link>
        <router-link to="/about">
          <li class="nav">
            <span>关于</span>
          </li>
        </router-link>
        <router-link to="/mainSshClient">
          <li class="nav">
            <span>服务器客户端</span>
          </li>
        </router-link>
      </ul>
    </nav>
    <div class="info">
      <img src="../assets/head.png" alt="head">
      <div class="info_name">吴亦堂</div>
      <div class="archive">
        <ul>
          <router-link to="/article">
            <li>
              <span class="archive_count">{{ articleNumber }}</span>
              <span class="archive_name">博客</span>
            </li>
          </router-link>
        </ul>
      </div>
      <ul class="communication">
        <li class="communication_item">
          <a href="https://blog.csdn.net/cxt520ppgsh">
            <span>吴亦堂的开发小站</span>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'sidebar',
    data() {
      return {
        articleNumber: ''
      }
    },
    mounted: function () {
      this.$http.post(global.articleList).then(
        response => {
          this.articleNumber = JSON.parse(response.data)["data"].length;
          console.log(this.articleNumber)
        },
        response => console.log(response)
      )
    }
  }
</script>

<style>
  #left_menu {
    float: left;
    width: 240px;
    margin-right: 20px;
  }

  @media (max-width: 767px) {
    #left_menu {
      display: none;
    }
  }

  #left_menu a:hover {
    color: #222;
    background: #f5f7f9;
  }

  .menu_name {
    background: rgb(38, 42, 48);
    color: white;
    padding: 10px 0;
  }

  nav {
    background: white;
    padding: 20px 0;
  }

  .nav {
    padding: 8px 20px;
    text-align: left;
  }

  .info {
    margin-top: 10px;
    padding: 20px;
    background: white;
  }

  .info_name {
    font-weight: 600;
    padding: 5px;
  }

  .info img {
    width: 120px;
    height: 120px;
  }

  .archive {
    padding-top: 20px;
  }

  .archive .archive_count {
    display: block;
    color: #222;
    font-weight: 600;
    font-size: 16px;
  }

  .archive .archive_name {
    display: block;
    color: #999;
    font-size: 14px;
  }

  .communication {
    padding-top: 20px;
  }

  .communication_item {
    display: inline-block;
    width: 40%;
    font-size: 14px;
    font-weight: 600;
    padding: 5px;
  }
</style>
