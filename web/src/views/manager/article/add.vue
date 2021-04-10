<template>
    <div class="article-add">
        <div class="editor">
            <vue-scroll :ops="ops">
                <div class="wrap">
                    <h1>{{$route.params.id?'更新文章':'新建文章'}}</h1>
                    <el-form label-width="80px">
                        <el-form-item label="文章标题">
                            <el-input palceholder="文章标题" v-model="param.title"></el-input>
                        </el-form-item>
                        <el-form-item label="分类">
                            <el-select v-model="param.category">
                                <el-option v-for="item in categories" :key="item.id" :value="item.id"
                                           :label="item.name"></el-option>
                            </el-select>
                            <el-button style="margin-left: 20px;" type="text">添加分类</el-button>
                        </el-form-item>
                        <el-form-item label="标签">
                            <el-checkbox-group v-model="param.tags">
                                <el-checkbox v-for="item in tags" :key="item.id" :label="item.id">{{item.name}}
                                </el-checkbox>
                            </el-checkbox-group>
                        </el-form-item>
                        <el-form-item label="文章封面">
                            <upload :file="param.thumb" :file-type="2" @success="handleUploadSuccess"
                                    @remove="handleRemove"/>
                            <el-input style="width: 200px;" v-model="param.thumb"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <mavon-editor v-model="param.content" @imgAdd="handleImgAdd"/>
                        </el-form-item>
                    </el-form>
                </div>
            </vue-scroll>
        </div>
        <div class="footer">
            <el-button type="default" @click="$router.go(-1)">取消</el-button>
            <el-button type="primary" block @click="handleAddArticle">确定</el-button>
        </div>
    </div>
</template>

<script>
import { AddArticle, GetArticle, UpdateArticle } from '@/api/article'
import MavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import Upload from '@/components/upload/image'
import { scroll, BASE_URL } from '@/config'
import { GetAllTags } from '@/api/tags'
import { GetAllCategories } from '@/api/category'
import axios from 'axios'
export default {
  components: {
    Upload,
    MavonEditor: MavonEditor.mavonEditor
  },
  data () {
    return {
      ops: scroll,
      tags: [],
      categories: [],
      param: {
        title: '',
        thumb: '',
        categoryId: 1,
        tags: [],
        content: ''
      }
    }
  },
  methods: {
    init () {
      GetAllTags().then(res => {
        if (res.code === 0) {
          this.tags = res.data || []
        }
      })
      GetAllCategories().then(res => {
        if (res.code === 0) {
          this.categories = res.data || []
        }
      })
      if (this.$route.params.id) {
        GetArticle({ id: this.$route.params.id }).then(res => {
          if (res.code === 0) {
            let param = res.data
            this.param = {
              ...param,
              category: param.categoryId,
              tags: param.tags.map(v => v.id)
            }
          }
        })
      }
    },
    handleUploadSuccess (url) {
      this.param.thumb = url
    },
    handleRemove () {
      this.param.thumb = ''
    },
    handleAddArticle () {
      if (!this.$route.params.id) {
        AddArticle(this.param).then(res => {
          if (res.code === 0) {
            this.$message.success('新建成功')
          }
        })
      } else {
        UpdateArticle(this.param).then(res => {
          if (res.code === 0) {
            this.$message.success('更新成功')
          }
        })
      }
    },
    handleImgAdd (pos, file) {
      let formData = new FormData()
      formData.append('type', 2)
      formData.append('file', file)
      axios({
        url: BASE_URL + 'api/file/upload',
        method: 'post',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        data: formData
      }).then(res => {
        if (res.data.code === 0) {
          let content = this.param.content
          let oStr = `(${pos})`
          let nStr = `(${res.data.data})`
          let index = content.indexOf(oStr)
          let str = content.replace(oStr, '')
          let insertStr = (soure, start, newStr) => {
            return soure.slice(0, start) + newStr + soure.slice(start)
          }
          this.param.content = insertStr(str, index, nStr)
        }
      })
    }
  },
  created () {
    this.init()
  }
}
</script>

<style lang="scss" scoped>
    .article-add {
        height: 100%;
        display: flex;
        background-color: #fff;
        flex-direction: column;

        .editor {
            flex: 1;
            overflow: hidden;
        }

        .footer {
            padding: $-padding;
            text-align: center;
            border-top: 1px solid $-bgc;
        }

        .wrap {
            padding: 0 50px;
        }
    }
</style>
