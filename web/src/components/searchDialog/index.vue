<template>
    <transition name="el-zoom-in-top">
    <div class="search" v-show="visible">
        <div class="form-item">
            <el-input v-model="param.keyword" placeholder="文章标题"></el-input>
        </div>
        <div class="form-item">
            <span class="label">分类：</span>
            <span class="checked" @click="param.category = 0" :class="{active: param.category===0}">全部</span>
            <span class="checked" @click="param.category = item.id" :class="{active: param.category === item.id}"  v-for="item in allCategories" :key="item.id">{{item.name}}</span>
        </div>
        <div class="form-item">
            <span class="label">标签：</span>
            <span class="checked" @click="handleCheckTag(item)" :class="{active: param.tags.includes(item.id)}"  v-for="item in allTags" :key="item.id">{{item.name}}</span>
        </div>
        <div class="footer">
            <el-button type="text" @click="handleClear">重置查询条件</el-button>
            <el-button type="primary" @click="handleQuery">查询</el-button>
        </div>
    </div>
    </transition>
</template>

<script>
export default {
  props: {
    visible: Boolean,
    tags: {
      type: Array,
      default: () => []
    },
    categories: {
      type: Array,
      default: () => []
    }
  },
  computed: {
    allTags () {
      return this.tags || []
    },
    allCategories () {
      return this.categories || []
    }
  },
  data () {
    return {
      param: {
        keyword: '',
        category: 0,
        tags: []
      }
    }
  },
  methods: {
    handleCheckTag (item) {
      let i = this.param.tags.indexOf(item.id)
      if (i >= 0) {
        this.param.tags.splice(i, 1)
      } else {
        this.param.tags.push(item.id)
      }
    },
    handleClear () {
      this.param = {
        keyword: '',
        category: 0,
        tags: []
      }
    },
    handleQuery () {
      this.$emit('search', { ...this.param })
    }
  }
}
</script>

<style lang="scss" scoped>
    .search{
        width: 100%;
        background: #fff;
        position: fixed;
        z-index: 10;
        top: 60px;
        box-shadow: inset 0 5px 20px rgba(0,0,0,.2);
        padding: $-padding;
    }
    .form-item{
        line-height: 30px;
        font-size: 14px;
        color: #666;
        vertical-align: middle;
        &:first-child{
            margin-bottom: 15px;
        }
        .label{
            font-weight: 700;
        }
        .checked{
            line-height: 1;
            display: inline-block;
            padding: 5px 10px;
            border: 1px solid transparent;
            font-size: 16px;
            margin-right: 15px;
            margin-bottom: 15px;
            &.active{
                color: $-active-color;
                border-color: $-active-color;
            }
        }
    }
    .footer{
        text-align: right;
    }
</style>
