<template>
    <div class="personal">
            <div class="wrap">
                <div class="txt">
                    <p class="name">JACOB WOO</p>
                    <p class="tip">never say never</p>
                </div>
                <div class="scroll-text">
                    <span>SCROLL</span>
                    <i class="el-icon-arrow-down"></i>
                </div>
            </div>
            <h2 class="sub-title">
               <span>TOP</span>
            </h2>
            <div class="article-wrap">
                <template v-for="item in articles">
                    <article-item :item="item" :key="item.id"/>
                </template>
            </div>
    </div>
</template>

<script>
import { GetArticleList } from '@/api/article'
import ArticleItem from '@/components/articleItem'
export default {
  data () {
    return {
      articles: []
    }
  },
  methods: {
    init () {
      GetArticleList({ pageNo: 1, pageSize: 5 }).then(res => {
        if (res.code === 0) {
          this.articles = res.data || []
        }
      })
    }
  },
  mounted () {
    this.init()
  },
  components: {
    ArticleItem
  }
}
</script>

<style lang="scss" scoped>
    @keyframes txtAnimate {
        from {
            height: 0;
        }
        to{
            height: 100vh;
        }
    }
    @keyframes scroll {
        from {
            transform: translateY(0px);
            opacity: 1;
        }
        to {
            transform: translateY(15px);
            opacity: 0;
        }
    }
    .personal{
        width: 100%;
        background-color: #fff;
        .wrap{
            height: 100vh;
            background-color: $-theme;
            background-image: url("../../assets/img/banner.jpg");
            background-attachment: fixed;
            background-repeat: no-repeat;
            background-position: center center;
            background-size: cover;
            position: relative;
            display: flex;
            justify-content: space-around;
            align-items: center;
            &:before{
                position: absolute;
                content: "";
                width: 100%;
                height: 100%;
                top: 0px;
                left: 0px;
                background-image: url("../../assets/img/spllit.png");
                background-repeat: repeat;
                background-attachment: scroll;
                opacity: 0.5;
                z-index: 1;
            }
            .txt{
                position: relative;
                z-index: 1;
                color: #fff;
                font-family: Serif;
                font-size: 32px;
                text-align: center;
                .name{
                    font-family: Serif;
                    font-weight: 700;
                }
                .tip{
                    font-size: 14px;
                    color: #888;
                }
            }
            .scroll-text{
                position: absolute;
                bottom: 20px;
                left: 50%;
                transform: translateX(-50%);
                color: #fff;
                z-index: 2;
                font-family: 'Merriweather', serif;
                letter-spacing: 2px;
                font-size: 18px;
                text-align: center;
                span{
                    display: block;
                }
                i{
                    animation: scroll 1s linear infinite;
                }
            }
        }
    }
    .article-wrap{
        width: 100%;
        padding: $-padding 0;
        overflow: hidden;
    }
    .sub-title{
        font-size: 22px;
        text-align: center;
        margin-top: 60px;
        position: relative;
        color: #333;
        &:after{
            content: '';
            display: block;
            width: 200px;
            height: 1px;
            background-color: #333;
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
        }
        span{
            padding: 10px;
            background-color: #fff;
            position: relative;
            z-index: 1;
        }
    }
</style>
