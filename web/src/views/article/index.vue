<template>
    <div class="article">

        <div class="left">
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
        <div class="right">
            <div class="title">
                <span class="blog-category"></span>文章分类
            </div>
            <div class="content">
                <div class="category-item" :class="{active: pagination.category === 0}"
                     @click="handleQueryByCategory(0)">全部
                </div>
                <div class="category-item" :class="{active: pagination.category === item.id}" v-for="item in categories"
                     :key="item.id" @click="handleQueryByCategory(item.id)">{{item.name}}
                </div>
            </div>
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
      colors: ['EF476F', 'ebad1d', '06D6A0', '118AB2', '073B4C'],
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
      loadCompleted: false
    }
  },
  methods: {
    init () {
      GetArticleUnion(this.pagination).then(res => {
        if (res.code === 0) {
          this.articles = res.data.list || []
          this.pagination.total = res.data.total
        }
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
      GetArticleUnion(this.pagination).then(res => {
        if (res.code === 0) {
          let data = res.data.list || []
          if (data.length === 0) {
            this.loadCompleted = true
          }
          this.articles = [...this.articles, ...data]
        }
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
      if (process === 1 && !this.loadCompleted) {
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
            width: 250px;
            background-color: #fff;

            .title {
                height: 50px;
                line-height: 50px;
                padding-left: $-padding;
                font-size: 16px;
                color: #7bb360;
                position: relative;

                span {
                    margin-right: $-margin;
                }

                &:after {
                    content: "";
                    display: block;
                    height: 10px;
                    width: 150px;
                    background: #7bb360;
                    opacity: .1;
                    position: absolute;
                    bottom: 14px;
                }

                &.tag-title {
                    color: #f56c6c;

                    &:after {
                        background: #f56c6c;
                    }
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
        line-height: 30px;
        margin-right: $-margin*2;
        margin-bottom: $-margin*2;
        font-size: 16px;
        cursor: pointer;
        color: #666;

        &:hover {
            text-decoration: underline;
        }

        &.active {
            color: $-active-color;
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
    }
</style>
