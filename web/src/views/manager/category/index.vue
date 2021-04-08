<template>
    <div class="categories">
        <vue-scroll :ops="ops">
            <div class="category-wrap">
                <h1>
                    文章分类管理
                </h1>
                <div>
                    <el-button type="primary" @click="handleAddCategory">添加分类</el-button>
                </div>
                <el-table :data="categories" style="width: 500px;">
                    <el-table-column label="序号" type="index" width="100px"></el-table-column>
                    <el-table-column label="分类名称" prop="name"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template slot-scope="scope">
                            <el-button type="text" @click="handleDel(scope.row, scope.$index)">删除</el-button>
                            <el-button type="text" @click="handleUpdate(scope.row)">修改</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </vue-scroll>
    </div>
</template>

<script>
import { GetAllCategories, AddCategory, UpdateCategory, DelCategory } from '@/api/category'
import { scroll } from '@/config'

export default {
  data () {
    return {
      ops: scroll,
      categories: []
    }
  },
  methods: {
    init () {
      GetAllCategories().then(res => {
        if (res.code === 0) {
          this.categories = res.data || []
        }
      })
    },
    handleAddCategory () {
      this.$prompt('请输入分类名称', '温馨提示').then(({ value }) => {
        AddCategory({ name: value }).then(res => {
          if (res.code === 0) {
            this.init()
            this.$message.success('添加成功')
          }
        })
      })
    },
    handleUpdate (item) {
      this.$prompt('请输入分类名称', '温馨提示', {
        inputValue: item.name
      }).then(({ value }) => {
        UpdateCategory({
          id: item.id,
          name: value
        }).then(res => {
          if (res.code === 0) {
            this.init()
            this.$message.success('添加成功')
          }
        })
      })
    },
    handleDel (item, index) {
      console.log(index)
      this.$confirm('是否确认删除', '温馨提示').then(() => {
        return DelCategory({ id: item.id })
      }).then(res => {
        if (res.code === 0) {
          this.$message.success('删除成功')
          this.categories.splice(index, 1)
        }
      })
    }
  },
  mounted () {
    this.init()
  }
}
</script>

<style lang="scss" scoped>
    .categories {
        height: 100%;
        background-color: #fff;
        overflow: hidden;

        .category-wrap {
            padding: $-padding*2;
        }
    }
</style>
