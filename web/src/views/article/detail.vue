<template>
    <div class="article-detail" v-loading="loading">
        <div class="left">
            <vue-scroll :ops="ops" ref="articleScroll">
            <header v-cloak>
                    <h1>
                    <span class="category">
                        【{{param.category && param.category.name}}】
                    </span>
                    {{param.title}}</h1>
                    <div class="tag" v-for="tag in param.tags" :key="tag.id"
                         :style="{color: '#fff', borderColor: tag.color, background:tag.color}">{{tag.name}}
                    </div>
                    <span class="label">
                        <i class="blog-view"></i>
                        {{param.view}}
                    </span>
                    <span class="label like">
                        <i class="blog-like icon" @click="handleLike"></i>
                        {{param.likes}}
                    </span>
                    <div class="date">
                        <span>创建：{{param.createdAt | formatDateTime}}</span>
                        <span>更新：{{param.updatedAt | formatDateTime}}</span>
                    </div>
                    <div class="close" @click="$router.go(-1)">
                        <i class="el-icon-close"></i>
                    </div>
            </header>
            <div class="article">
                    <div class="article-content">
                        <div v-html="param.content">
                        </div>
                    </div>
            </div>
            <comment :article-id="param.id"></comment>
            </vue-scroll>
        </div>
        <div class="right">
            <vue-scroll :ops="ops">
                <div class="article-toc-content" @click.stop.prevent="handleAnchor"></div>
            </vue-scroll>
        </div>
    </div>
</template>

<script>
import { GetArticle, like } from '@/api/article'
import { translate } from '@/utils'
import { scroll } from '@/config'
import tocbot from 'tocbot'
import Comment from './comment'

export default {
  data () {
    return {
      loading: false,
      ops: scroll,
      param: {
        title: '',
        content: ''
      }
    }
  },
  methods: {
    init () {
      this.loading = true
      GetArticle({ id: this.$route.params.id }).then(res => {
        if (res.code === 0) {
          let param = res.data
          param.content = translate(param.content)
          this.param = param
          this.$nextTick(() => {
            tocbot.init({
              tocSelector: '.article-toc-content',
              contentSelector: '.article-content',
              headingSelector: 'h1, h2, h3, h4, h5, h6',
              hasInnerContainers: true
            })
          })
        }
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleAnchor (e) {
      if (e.target.tagName === 'A') {
        let target = e.target.href
        let id = target.split('#')[1]
        this.$refs.articleScroll.scrollIntoView(`#${decodeURI(id)}`, 300)
      }
      // if(e.target === '')
    },
    handleLike () {
      like(this.param.id).then(res => {
        if (res.code === 0) {
          this.param.likes += 1
        }
      })
    }
  },
  mounted () {
    this.init()
  },
  components: {
    Comment
  }
}
</script>

<style lang="scss">
    .article-detail {
        height: 100vh;
        padding-right: 0;
        display: flex;
        overflow: hidden;

        .left {
            flex: 1;
            height: 100%;
            overflow: hidden;
            border-right: 1px solid $-bgc;

            header {
                height: 100px;
                background-color: #fff;
                border-bottom: 1px solid $-bgc;
                padding: $-padding;
                color: #333;
                position: relative;
                .label{
                    font-size: 16px;
                    color: #d9e1e7;
                    margin-left: $-margin*2;
                    &.like{
                        cursor: pointer;
                        .icon{
                            color: $-red;
                        }
                    }
                }
                .date{
                    float: right;
                    color: #80929f;
                    span{
                        margin-right: $-margin * 2;
                    }
                }
            }

            .article {
                flex: 1;
                overflow: hidden;
            }
        }

        .right {
            width: 250px;
            height: 100%;
            overflow: hidden;
            background-color: #fff;

            .article-toc-content {
                ol {
                    padding-left: 20px;
                    margin-bottom: 10px;
                }

                li {
                    line-height: 30px;
                }

            }
        }

        a {
            color: #3a8ee6;

            &:hover {
                color: #2e7ae6;
            }
        }

        .close {
            font-size: 24px;
            font-weight: 700;
            color: #bbbbbb;
            position: absolute;
            right: 10px;
            top: 10px;
            cursor: pointer;
            &:hover{
               color: #999;
            }
        }
    }

    .article-content {
        height: 100%;
        background: #fff;
        padding: $-padding*2;
        margin: 0 auto;
        font-family: 'Lato', 'PingFang SC', 'Microsoft YaHei', sans-serif;
        color: #555;
        line-height: 2;
        font-size: 14px;

        img {
            max-width: 100%;
            margin: 0 auto 25px;
            box-sizing: border-box;
            padding: 3px;
            border: 1px solid #ddd;
        }

        pre,
        code {
            font-family: consolas, Menlo, 'PingFang SC', 'Microsoft YaHei', monospace;
        }

        code {
            padding: 2px 4px;
            word-wrap: break-word;
            color: #ff502c;
            background: #fff5f5;
            border-radius: 3px;
            font-size: 13px;
        }

        pre {
            padding: 10px;
            overflow: auto;
            margin: 20px 0;
            font-size: 13px;
            color: #4d4d4c;
            background: #f7f7f7;
            line-height: 1.6;
        }

        pre code {
            padding: 0;
            color: #555;
            background: none;
            text-shadow: none;
            font-family: consolas, Menlo, 'PingFang SC', 'Microsoft YaHei', monospace;
        }

        h2,
        h3,
        h4,
        h5,
        h6 {
            margin: 20px 0 15px 0;
            padding-top: 10px;
            border-bottom: 1px solid #eee;
            margin-bottom: 10px;
            font-weight: bold;
            line-height: 1.5;
            font-family: 'Lato', 'PingFang SC', 'Microsoft YaHei', sans-serif;
            color: #555;
        }

        h2 {
            font-size: 1.4em;
        }

        h3 {
            font-size: 1.3em;
            border-bottom: 1px dotted #eee;
        }

        h4 {
            font-size: 1.2em;
        }

        ul {
            padding-left: 20px;

            li {
                line-height: 2;
                list-style: circle;
            }
        }

        .hljs-comment,
        .hljs-quote {
            color: #575f6a;
        }

        blockquote {
            margin: 1em 0;
            border-left: 4px solid #ddd;
            padding: 0 1em;
            color: #666;

            p {
                margin: .5em 0;
                line-height: 1.7em;
            }
        }

        table {
            font-size: 1em;
            max-width: 100%;
            overflow: auto;
            border: 1px solid #f6f6f6;
            border-collapse: collapse;
            border-spacing: 0;

            thead {
                background: #f6f6f6;
                color: #000;
                text-align: left;
            }

            tr {
                display: table-row;
                vertical-align: inherit;
                border-color: inherit;

                &:nth-child(2n) {
                    background-color: #fcfcfc;
                }
            }

            th {
                padding: 0.8em 0.5em;
                line-height: 2em;
            }

            tbody {
                display: table-row-group;
                vertical-align: middle;
                border-color: inherit;

                td {
                    min-width: 7.5em;
                    padding: 0.8em 0.5em;
                    line-height: 1.5em;
                }
            }
        }
    }

    @media (max-width: 780px) {
        .article-detail {
            padding-top: 60px;
            .left{
                header{
                    height: 80px;
                    h1{
                        font-size: 18px;
                    }
                    .tag{
                        display: none;
                        padding: 5px;
                        font-size: 12px;
                    }
                    .label{
                        display: none;
                    }
                    .date{
                        float: none;
                    }
                }
            }
        }
        .right{
            display: none;
        }

    }

</style>
