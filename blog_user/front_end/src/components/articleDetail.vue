<template>
      <div class="aligin_left" v-html="compiledMarkdown()"></div>
</template>

<script>
    import marked from 'marked'
    import hljs from "highlight.js";
    import '../assets/atom-one-light.css'
    marked.setOptions({
            renderer: new marked.Renderer(),
            highlight: function(code) {
                return hljs.highlightAuto(code).value;
            },
            pedantic: false,
            gfm: true,
            tables: true,
            breaks: false,
            sanitize: false,
            smartLists: true,
            smartypants: false,
            xhtml: false
        }
    );
    export default {
        name: 'articleDetail',
        data() {
            return {
                article: {}
            }
        },
        mounted: function () {
            window.scrollTo(0,0);
            let _articleId = this.$route.params.articleId
            this.$axios.post(global.articleDetail,{
                'articleId' : _articleId
            }).then(
                response => {
                    this.article = JSON.parse(response.data)["data"]
                },
                response => console.log(response)
            )
        },
        methods: {
            compiledMarkdown: function () {
                console.log(this.article.content)
                return marked(this.article.content || '', {sanitize: true})
            }
        }
    }
</script>

<style>
  .aligin_left {
    text-align: left;
    float: left;
    padding: 0;
    font-size: 16px;
    overflow-x: scroll;
    width: 60%;
  }
</style>
