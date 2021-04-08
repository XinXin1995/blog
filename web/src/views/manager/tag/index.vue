<template>
    <div class="tags">
        <h1>
            文章标签管理
            <el-button type="primary" @click="handleAddTag">新增标签</el-button>
        </h1>
        <div class="tag closable"
             v-for="(item,index) in tags"
             :key="item.id"
             :color="item.color"
             :style="{
             color: '#fff',
             backgroundColor: item.color,
             borderColor: item.color,
             }"
        >
            {{item.name}}
            <i class="el-icon-edit" @click="handleUpdateTag(item)"></i>
            <i class="el-icon-close" @click="handleDelTag(item,index)"></i>
        </div>
        <el-dialog width="400px" :before-close="handleBeforeClose" :visible.sync="visible" :title="param.id?'新增标签':'修改标签'">
            <el-form label-width="60px" :model="param" :rules="rules">
                <el-form-item label="名称" prop="name">
                    <el-input v-model="param.name"></el-input>
                </el-form-item>
                <el-form-item label="颜色" prop="color">
                    <el-color-picker v-model="param.color"/>
                </el-form-item>
            </el-form>
            <div slot="footer">
                <el-button @click="handleClose">取消</el-button>
                <el-button type="primary" @click="handleConfirm">确定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import { GetAllTags, DelTag, AddTag, UpdateTag } from '@/api/tags'

export default {
  data () {
    return {
      visible: false,
      param: {
        name: '',
        color: ''
      },
      tags: [],
      rules: {
        name: [
          { required: true, message: '请输入标签名称', trigger: 'blur' }
        ],
        color: [
          { required: true, message: '请选择颜色', trigger: 'blur' }
        ]
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
    },
    handleDelTag (item, index) {
      this.$confirm('是否确认删除？', '温馨提示').then(() => {
        return DelTag({ id: item.id })
      }).then(res => {
        if (res.code === 0) {
          this.tags.splice(index, 1)
        }
      })
    },
    handleBeforeClose (done) {
      this.handleClose()
      done()
    },
    handleClose () {
      this.param = {
        name: '',
        color: ''
      }
      this.visible = false
    },
    handleAddTag () {
      this.param = {
        name: '',
        color: ''
      }
      this.visible = true
    },
    handleUpdateTag (item) {
      this.param = { ...item }
      this.visible = true
    },
    handleConfirm () {
      // eslint-disable-next-line func-call-spacing
      const $ = this
      if ($.param.id) {
        UpdateTag($.param).then(res => {
          if (res.code === 0) {
            this.handleClose()
            this.init()
          }
        })
      } else {
        AddTag($.param).then(res => {
          if (res.code === 0) {
            this.handleClose()
            this.init()
          }
        })
      }
    }
  },
  mounted () {
    this.init()
  }
}
</script>

<style lang="scss" scoped>
    .tags {
        height: 100%;
        background-color: #fff;
        overflow: hidden;
        padding: $-padding * 3;
        .tag{
            &.closable{
                i{
                    cursor: pointer;
                    margin-left: 10px;
                }
            }
        }
    }
</style>
