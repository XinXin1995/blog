<template>
    <div class="article">
        <div class="left" v-show="!loading" v-loading="pullLoading">
            <span class="search-btn" @click="searchVisible = !searchVisible">
                <i class="el-icon-search"></i>
            </span>
            <search-dialog :visible="searchVisible" :tags="tags" :categories="categories" @search="handleSearch"/>
            <vue-scroll ref="scroll" :ops="ops" @handle-scroll="handleLoadMore">
                <div class="wrap">
                    <template v-for="item in articles">
                        <article-item :item="item" :key="item.id" @like="handleLike(item)"/>
                    </template>
                    <div class="load-complete-text" v-if="loadCompleted">
                        已加载全部
                    </div>
                </div>
            </vue-scroll>
        </div>
        <div class="skeleton" v-show="loading">
            <div class="list">
                <div class="skeleton-article-item" v-for="i in 10" :key="i">
                    <div class="text">
                        <h1 class="skeleton-animation"></h1>
                        <div class="tags">
                            <span class="skeleton-animation"></span>
                            <span class="skeleton-animation"></span>
                        </div>
                    </div>
                    <div class="img skeleton-animation"></div>
                </div>
            </div>
        </div>
        <div class="right">
            <div class="card">
                <div class="title">
                    <span class="blog-category"></span>文章分类
                </div>
                <div class="content">
                    <div class="category-item"
                         @click="handleQueryByCategory(0)">
                        <div :style="{
                            backgroundColor: pagination.category === 0 ? '#fff':'#6890da',
                            borderColor: '#6890da',
                            color: pagination.category === 0 ? '#6890da' : '#f5f5f5'
                          }"
                        >全部</div>
                    </div>
                    <div class="category-item"
                         v-for="(item, index) in categories"
                         :key="item.id" @click="handleQueryByCategory(item.id)">
                        <div :style="{
                        backgroundColor: pagination.category === item.id? '#fff':colors[(index + 5)%5],
                        borderColor:  colors[(index + 5)%5],
                        color: pagination.category === item.id ? colors[(index + 5)%5] : '#f5f5f5'
                        }">
                            {{item.name}}
                        </div>
                    </div>
                </div>
            </div>

            <div class="card">
            <div class="title tag-title">
                <span class="blog-tag"></span>文章标签
            </div>
            <div class="content">
                <div class="tag" v-for="item in tags"
                     @click="handleQueryByTag(item)"
                     :class="{active: pagination.tags.includes(item.id)}"
                     :style="{color: '#fff',backgroundColor: item.color, borderColor: item.color}"
                     :key="item.id">{{item.name}}
                </div>
            </div>
            </div>
        </div>
    </div>
</template>

<script>
import { GetArticleUnion } from '@/api/article'
import { GetAllTags } from '@/api/tags'
import { GetAllCategories } from '@/api/category'
import { scroll } from '@/config'
import ArticleItem from '@/components/articleItem'
import SearchDialog from '@/components/searchDialog'

export default {
  components: {
    ArticleItem,
    SearchDialog
  },
  data () {
    return {
      ops: scroll,
      colors: ['#da5252', '#e3974a', '#47be66', '#43a2c2', '#3e6674'],
      tags: [],
      categories: [],
      articles: [],
      pagination: {
        pageNo: 1,
        pageSize: 10,
        keyword: '',
        total: 0,
        category: 0,
        tags: []
      },
      searchVisible: false,
      loadCompleted: false,
      loading: false,
      pullLoading: false
    }
  },
  methods: {
    init () {
      this.loading = true
      GetArticleUnion(this.pagination).then(res => {
        if (res.code === 0) {
          this.articles = res.data.list || []
          this.pagination.total = res.data.total
        }
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    initTags () {
      GetAllTags().then(res => {
        if (res.code === 0) {
          this.tags = res.data || []
        }
      })
    },
    initCategories () {
      GetAllCategories().then(res => {
        if (res.code === 0) {
          this.categories = res.data || []
        }
      })
    },
    initLoadMore () {
      this.pullLoading = true
      GetArticleUnion(this.pagination).then(res => {
        if (res.code === 0) {
          let data = res.data.list || []
          if (data.length === 0) {
            this.loadCompleted = true
          }
          this.articles = [...this.articles, ...data]
        }
        this.pullLoading = false
      }).catch(() => {
        this.pullLoading = false
      })
    },
    handleLike (item) {
      item.likes = item.likes + 1
    },
    handleToDetail (item) {
      this.$router.push({ name: 'article.detail', params: { id: item.id } })
    },
    handleQueryByTag (item) {
      this.pagination.pageNo = 1
      let i = this.pagination.tags.indexOf(item.id)
      if (i >= 0) {
        this.pagination.tags.splice(i, 1)
      } else {
        this.pagination.tags.push(item.id)
      }
      this.pagination.tags.sort()
      this.init()
      this.$refs.scroll.scrollTo({ y: 0, animation: 0 })
      this.loadCompleted = false
    },
    handleQueryByCategory (category) {
      this.pagination.pageNo = 1
      this.pagination.category = category
      this.$refs.scroll.scrollTo({ y: 0, animation: 0 })
      this.init()
      this.loadCompleted = false
    },
    handleLoadMore ({ process }) {
      // console.log(e)
      if (process >= 0.8 && !this.loadCompleted) {
        this.pagination.pageNo += 1
        this.initLoadMore()
      }
    },
    handleSearch (param) {
      this.pagination.tags = param.tags
      this.pagination.category = param.category
      this.pagination.keyword = param.keyword
      this.searchVisible = false
      this.init()
    }
  },
  mounted () {
    this.init()
    this.initTags()
    this.initCategories()
  }
}
</script>

<style lang="scss" scoped>
    .article {
        height: 100vh;
        display: flex;

        .right {
            height: 100%;
            width: 400px;
            padding: $-padding;
            .card{
                background: #fff;
                border: 1px solid #e8e8e8;
                & + .card{
                    margin-top: 24px;
                }
            }
            .title {
                height: 50px;
                line-height: 50px;
                padding-left: $-padding;
                font-size: 16px;
                color: #666;
                position: relative;

                span {
                    margin-right: $-margin;
                }
            }

            .content {
                padding: $-padding;
            }

            .tag {
                &.active {
                    position: relative;

                    &:after {
                        content: '';
                        display: block;
                        width: 20px;
                        height: 20px;
                        background-color: #fff;
                        border-radius: 50%;
                        position: absolute;
                        top: -10px;
                        right: -10px;
                    }
                }
            }
        }

        .left {
            flex: 1;
            overflow: hidden;
        }
    }

    .skeleton {
        height: 100%;
        flex: 1;
        display: flex;
        overflow: hidden;

        .list {
            flex: 1;
            overflow: hidden;
            padding: $-padding*2 0;
        }

        &-article-item {
            max-width: 920px;
            width: 90%;
            margin: 0 auto;
            margin-bottom: $-margin * 2;
            background-color: #fff;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, .1);
            height: 230px;
            overflow: hidden;
            display: flex;

            &:nth-child(2n) {
                flex-direction: row-reverse;

                .text {
                    & > * {
                        float: right;
                    }

                    .tags {
                        span {
                            margin-right: 0;
                            margin-left: $-margin;
                        }
                    }
                }
            }

            .text {
                flex: 1;
                padding: 0 $-padding;

                h1 {
                    height: 32px;
                    width: 60%;
                }

                .tags {
                    span {
                        display: inline-block;
                        width: 100px;
                        height: 32px;
                        margin-right: $-margin;
                    }
                }
            }

            .img {
                width: 50%;
            }
        }
    }

    .search-btn {
        position: fixed;
        width: 60px;
        height: 60px;
        color: #fff;
        font-size: 24px;
        top: 0;
        right: 0;
        z-index: 11;
        text-align: center;
        line-height: 60px;
        display: none;
    }

    .category-item {
        display: inline-block;
        height: 50px;
        line-height: 38px;
        font-size: 16px;
        cursor: pointer;
        width: 50%;
        text-align: center;
        color: #f2f4f5;

        &:nth-child(2n+1) {
            border-right: 12px solid #fff;
            border-bottom: 12px solid #fff;
        }

        &:nth-child(2n) {
            border-left: 12px solid #fff;
            border-bottom: 12px solid #fff;
        }

        & > div {
            width: 100%;
            height: 100%;
            border: 1px solid #666;
            border-radius: 5px;
            overflow: hidden;
        }
    }

    .left {
        height: 100%;

        .wrap {
            width: 100%;
            height: 100%;
            padding: $-padding*2 0;
        }
    }

    .load-complete-text {
        font-size: 16px;
        color: #999;
        text-align: center;
        line-height: 32px;
    }

    @keyframes skeleton-animate {
        0% {
            background-position: 100% 50%
        }
        to {
            background-position: 0 50%
        }
    }

    .skeleton-animation {
        background: linear-gradient(45deg, hsla(0, 0%, 74.5%, .2) 25%, hsla(0, 0%, 50.6%, .24) 37%, hsla(0, 0%, 74.5%, .2) 63%);
        background-size: 400% 100%;
        animation: skeleton-animate 1s ease infinite;
    }

    @media (max-width: 780px) {
        .article {
            margin-top: 60px;
            display: block;

            .right {
                display: none;
            }
        }
        .search-btn {
            display: block;
        }
        .skeleton-article-item {
            display: block;
            height: 300px;

            &:nth-child(2n) {
                .text > * {
                    float: none;
                }
            }

            .text {
                height: 150px;
            }

            .img {
                width: 100%;
                height: 150px;
            }
        }
    }
</style>
