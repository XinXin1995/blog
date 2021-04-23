<template>
    <div class="article-list">
        <vue-scroll :ops="ops">
            <div class="wrap">
                <h1>
                    文章列表
                    <el-button type="primary" @click="$router.push({name: 'manager.article.add'})">新建文章</el-button>
                </h1>
                <div class="utils">
                    <el-form :inline="true">
                        <el-form-item label="标题">
                            <el-input style="width: 200px;" v-model="pagination.keyword" placeholder="关键字查询"
                                      @change="init"></el-input>
                        </el-form-item>
                        <el-form-item label="分类">
                            <el-select v-model="pagination.category" @change="init">
                                <el-option label="全部" :value="0"></el-option>
                                <el-option v-for="item in categories" :key="item.id" :label="item.name"
                                           :value="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="标签">
                            <el-select multiple v-model="pagination.tags" style="width: 300px" @change="init">
                                <el-option v-for="item in tags" :key="item.id" :label="item.name"
                                           :value="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-form>
                </div>
                <el-table :data="articles" border>
                    <el-table-column align="center" width="100px" label="序号" type="index"></el-table-column>
                    <el-table-column label="标题">
                        <template slot-scope="scope">
                            <el-button type="text"
                                       @click="$router.push({name:'manager.article.update', params:{id:scope.row.id}})">
                                {{scope.row.title}}
                            </el-button>
                        </template>
                    </el-table-column>
                    <el-table-column align="center" width="150px" label="分类">
                        <template slot-scope="scope">
                            {{scope.row.categoryName || '--'}}
                        </template>
                    </el-table-column>
                    <el-table-column label="标签">
                        <template slot-scope="scope">
                            <div class="tag" v-for="tag in scope.row.tags" :key="tag.id"
                                 :style="{color: tag.color, borderColor: tag.color}">{{tag.name}}
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column align="center" width="100px" label="操作">
                        <template slot-scope="scope">
                            <el-button type="text" @click="handleDel(scope.row)"><i class="el-icon-delete"></i>
                            </el-button>
                            <el-button type="text" @click="handleTop(scope.row)"><i class="el-icon-caret-top"></i>
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
                <div class="pagination">
                    <el-pagination
                            @size-change="init"
                            @current-change="init"
                            :current-page.sync="pagination.pageNo"
                            :page-sizes="[10, 20, 30, 40]"
                            :page-size.sync="pagination.pageSize"
                            layout="sizes, prev, pager, next, jumper, ->, total"
                            :total="pagination.total">
                    </el-pagination>
                </div>
            </div>
        </vue-scroll>

    </div>
</template>

<script>
import { scroll } from '@/config'
import { GetArticleUnion, DelArticle, TopArticle } from '@/api/article'
import { GetAllTags } from '@/api/tags'
import { GetAllCategories } from '@/api/category'

export default {
  data () {
    return {
      ops: scroll,
      articles: [],
      categories: [],
      tags: [],
      pagination: {
        pageNo: 1,
        pageSize: 10,
        keyword: '',
        orderType: 3,
        tags: [],
        category: 0
      }
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
    handleDel (item) {
      this.$confirm('是否确认删除', '提示').then(() => {
        return DelArticle({ id: item.id })
      }).then(res => {
        if (res.code === 0) {
          this.$message.success('删除成功')
          this.init()
        }
      })
    },
    handleTop (item) {
      this.$confirm('是否确认置顶', '提示').then(() => {
        return TopArticle({ id: item.id })
      }).then(res => {
        if (res.code === 0) {
          this.$message.success('置顶成功')
          this.init()
        }
      })
    }
  },
  created () {
    this.init()
    GetAllTags().then(res => {
      this.tags = res.data || []
    })
    GetAllCategories().then(res => {
      this.categories = res.data || []
    })
  }
}
</script>

<style lang="scss" scoped>
    .article-list {
        overflow: hidden;
        height: 100%;
        background-color: #fff;

        .wrap {
            h1 {
                margin-top: 0;
            }

            padding: $-padding*3;
        }

        .pagination {
            padding: $-padding 0;
        }

        .utils {
            padding: $-padding 0;
        }
    }
</style>
