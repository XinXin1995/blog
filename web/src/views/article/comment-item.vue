<template>
    <div >
        <div class="comment-item" v-for="(item, index) in list" :key="index" @hover.stop>
            <div class="wrap">
                <div class="avatar">
                    <el-avatar :src="item.avatar" :size="40"></el-avatar>
                </div>
                <div class="text">
                    <div class="info">
                        <span>{{item.username}}</span>
                        {{item.createdAt | formatDateTime}}
                        <div class="btns">
                            <el-button type="text" @click="handleReply(item)">回复</el-button>
                        </div>
                    </div>
                    <p>
                        {{item.content}}
                    </p>
                </div>
            </div>
            <div class="children" v-if="item.children && item.children.length">
                <comment-item :list="item.children" @reply="handleReply"/>
            </div>
        </div>
    </div>
</template>

<script>

export default {
  name: 'CommentItem',
  props: {
    list: Array
  },
  data () {
    return {
      pagination: {
        pageNo: 1,
        pageSize: 10,
        articleId: 0
      }
    }
  },
  methods: {
    handleReply (item) {
      this.$emit('reply', item)
    }
  },
  mounted () {
  }
}
</script>

<style lang="scss" scoped>
    .comment-item {
        margin-bottom: $-margin*2;
    }
    .wrap {
        display: flex;
        padding: $-padding;
        &:hover{
            background-color: $-bgc;
        }

        .text {
            flex: 1;
            padding-left: $-padding;
            color: #999;

            .info {
                height: 32px;
                line-height: 32px;
                border-bottom: 1px solid $-bgc;
                span{
                    margin-right: 50px;
                }

                .btns {
                    float: right;

                    .el-button {
                        padding: 0;
                        line-height: 32px;
                        vertical-align: middle;
                    }
                }
            }
        }
    }
    .children {
        padding-left: 40px;
    }
</style>
