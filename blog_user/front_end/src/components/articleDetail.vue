<template>
  <div class="aligin_right"
       id="article">
    <markdown/>
  </div>
</template>

<script>
    import Markdown from '../assets/README.md'
    import 'highlight.js/styles/github-gist.css'

    export default {
        components: {Markdown}
    }
</script>

<style>
  .live-wrapper {
    background: #eee;
    padding: 10px;
  }

  .aligin_right {
    text-align: left;
    float: left;
    padding: 0;
    font-size: 22px;
  }
</style>
