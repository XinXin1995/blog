<template>
    <div class="comment">
        <h1 class="title">
            留言

            <span class="btns">
                <el-button type="text" @click="handleReply({})">留言</el-button>
            </span>
        </h1>
        <comment-item :list="comments" @reply="handleReply"/>
    </div>
</template>

<script>
import CommentItem from './comment-item'
import { AddComment, GetCommentsTree } from '@/api/comment'
import { Visit } from '@/api/user'
import { getToken, setToken } from '@/utils/auth'
import { mapActions } from 'vuex'

export default {
  props: {
    articleId: Number
  },
  data () {
    return {
      comments: []
    }
  },
  methods: {
    ...mapActions(['login']),
    init () {
      let articleId = Number(this.$route.params.id)
      GetCommentsTree({ articleId }).then(res => {
        if (res.code === 0) {
          this.comments = res.data || []
        }
      })
    },
    handleReply (item) {
      if (!getToken()) {
        this.$confirm('登录后才可以评论', '提示', {
          type: 'warning',
          confirmButtonText: '以游客方式登录'
        }).then(() => {
          return Visit({ username: '游  客', password: '123456' })
        }).then(res => {
          if (res.code === 0) {
            setToken(res.data.token)
          }
        }).catch(() => {

        })
        return
      }
      let articleId = this.articleId
      let parentId = item.id
      let title = parentId ? '请输入您的回复信息' : '请输入您的留言信息'
      this.$prompt(title, '提示', {
        inputType: 'textarea'
      }).then(res => {
        let content = res.value
        return AddComment({ articleId, parentId, content })
      }).then(res => {
        if (res.code === 0) {
          if (item && item.id) {
            item.children = item.children || []
            this.$set(item, 'children', [...item.children, res.data])
          } else {
            this.comments.unshift(res.data)
          }
        }
      }).catch(err => {
        console.error(err)
      })
    }
  },
  mounted () {
    this.init()
  },
  components: {
    CommentItem
  }
}
</script>

<style lang="scss" scoped>
    .comment {
        background: #fff;
        padding: 0 $-padding 100px;

        .title {
            margin: 0;
            height: 45px;
            line-height: 45px;
            border-bottom: 1px solid $-bgc;
            padding: 0 $-padding;
            margin-bottom: $-margin;

            .btns {
                float: right;
            }
        }
    }
</style>
