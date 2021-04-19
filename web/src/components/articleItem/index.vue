<template>
    <div class="article-item" @click="$router.push({name: 'article.detail', params: {id: item.id}})">
        <div class="detail">
            <h1>
                <span class="article-category">[{{item.category && item.category.name}}]</span>
                {{item.title}}
            </h1>
            <div class="article-tag">
                <div class="tag" v-for="it in item.tags" :key="`${item.id}-${it.id}`"
                     :style="{color: it.color,borderColor: it.color}">
                    {{it.name}}
                </div>
            </div>
            <div class="other" >
                <span><i class="blog-view"></i>{{item.view}}</span>
                <span class="like"><i class="blog-like" @click.stop.prevent="handleLike"></i>{{item.likes}}</span>
            </div>
            <div class="date">
                <span class="label">创建：{{item.createdAt | formatDateTime}}</span>
                <span class="label">更新：{{item.updatedAt | formatDateTime}}</span>
            </div>
        </div>
        <div class="img">
            <img v-lazy="item.thumb" :alt="item.title">
        </div>
    </div>
</template>

<script>
import { like } from '@/api/article'
export default {
  props: {
    item: {
      type: Object,
      default: () => ({})
    }
  },
  methods: {
    handleLike () {
      like(this.item.id).then(res => {
        if (res.code === 0) {
          this.$emit('like')
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
    .article-item  {
        max-width: 920px;
        width: 90%;
        margin: 0 auto;
        margin-bottom: $-margin * 2;
        background-color: #fff;
        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0,0,0,.1);
        height: 230px;
        overflow: hidden;
        display: flex;
        cursor: pointer;
        &:nth-child(2n){
            flex-direction: row-reverse;
            .detail{
                text-align: right;
            }
        }
        .detail{
            padding: $-padding*2;
            flex: 1;
            overflow: hidden;
        }
        .img{
            flex: 1;
            height: 100%;
            position: relative;
            overflow: hidden;
            display: flex;
            align-items: center;
            justify-content: space-around;
            background-color: $-bgc;
            img{
                width: 100%;
                min-height: 100%;
                transition: all .3s;
                &[lazy=error]{
                    width: auto;
                    height: auto;
                }
            }
            &:hover{
                img{
                    transform: scale(1.2);
                }
            }
        }

        h1 {
            margin-top: 0;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }

        .article-category {
            color: #999;
            font-weight: normal;
            cursor: pointer;

            &:hover {
                color: $-active-color;
            }
        }

        .article-tag {
            padding-left: 10px;
            .item {
                display: inline-block;
                padding: $-padding/2 $-padding;
                border: 1px solid #999;
                color: #999;
                border-radius: 5px;
                margin-right: $-margin;
            }
        }

        .other {
            height: 100px;
            padding-top: 35px;

            padding-left: 10px;
            span{
                color: #bcc6ce;
                margin-right: $-margin*3;
                font-size: 20px;
                i{
                    margin-right: $-margin;
                }
                &.like{
                    i{
                        color: $-red;
                    }
                }
            }
        }

        .date{
            padding-left: 10px;
            color: #bcc6ce;
            .label{
                margin-right: $-margin * 2;
            }
        }
    }
    @media (max-width: 780px) {
        .article-item {
            width: 90%;
            flex-direction: column-reverse !important;
            height: 300px;
            .img{
                flex: none;
                height: 180px;
            }
            .detail{
                flex: 1;
                text-align: left !important;
                overflow: hidden;
                padding: $-padding;
                h1{
                    font-size: 16px;
                }
                .other{
                    padding-top: 0;
                    height: 30px;
                }
                .article-tag{
                    padding-left: 0;
                    .tag{
                        margin-bottom: 5px;
                        font-size: 12px;
                    }
                }
                .other{
                    padding-left: 0;
                    line-height: 30px;
                    span{
                        font-size: 14px;
                    }
                }
                .date{
                    padding-left: 0;
                }
            }
        }
    }
</style>
