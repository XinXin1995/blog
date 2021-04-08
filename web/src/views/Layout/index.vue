<template>
    <div class="layout">
        <nav :class="{show:showNav}">
            <span class="menu-btn blog-menu" @click="menuVisible=true"></span>
            <logo></logo>
        </nav>
        <menu-dialog :visible.sync="menuVisible" :menu="menu"/>
        <aside>
            <logo></logo>
            <div class="avatar" v-if="user && user.avatar">
                <el-avatar :src="user.avatar" :size="120"></el-avatar>
            </div>
            <ul class="menu-list">
                <template v-for="(item, index) in menu">
                    <li :key="index" @click="handleDirect(item)" v-if="item.auth >= (user.role || 3)">
                        <span :class="item.icon"></span>
                        {{item.name}}
                    </li>
                </template>
            </ul>
            <auth/>
        </aside>
        <div class="submenu" v-if="submenu && submenu.length">
            <ul class="menu-list">
                <li v-for="(item, index) in submenu" :key="index" @click="handleDirectSub(item)">{{item.name}}</li>
            </ul>
        </div>
        <div class="content">
            <vue-scroll :ops="ops" @handle-scroll="handleScroll">
                <router-view></router-view>
            </vue-scroll>
        </div>
    </div>
</template>

<script>
import auth from './auth'
import { mapGetters } from 'vuex'
import { scroll } from '@/config'
import Logo from '@/components/logo'
import MenuDialog from '@/components/menu'
const menu = [
  {
    name: '首页',
    router: 'personal',
    icon: 'blog-home',
    auth: 3,
    submenu: []
  },
  {
    name: '博客',
    router: 'article.index',
    icon: 'blog-blog',
    auth: 3,
    submenu: []
  },
  {
    name: '后台',
    icon: 'blog-setting',
    router: 'manager',
    auth: 1,
    submenu: [
      {
        name: '文章管理',
        router: 'manager.article.list',
        auth: 1
      },
      {
        name: '文章分类',
        router: 'manager.category',
        auth: 1
      },
      {
        name: '文章标签',
        router: 'manager.tag',
        auth: 1
      }
    ]
  }
]
export default {
  computed: {
    ...mapGetters(['user']),
    inMobile () {
      const w = window.clientWidth || 780
      return w <= 780
    }
  },
  data () {
    return {
      menu,
      ops: scroll,
      submenu: [],
      showNav: false,
      menuVisible: false,
      searchVisible: false
    }
  },
  methods: {
    handleDirect (item) {
      if (item.router && item.router !== this.$route.name) {
        this.$router.push({ name: item.router })
      }
      this.submenu = item.submenu
    },
    handleDirectSub (item) {
      if (item.router && item.router !== this.$route.name) {
        this.$router.push({ name: item.router })
      }
    },
    handleScroll ({ process }) {
      if (process > 0 && this.$route.name === 'personal') {
        this.showNav = true
      }
    }
  },
  components: {
    auth,
    Logo,
    MenuDialog
  },
  created () {
    if (this.$route.name !== 'personal') {
      this.showNav = true
    }
  },
  watch: {
    $route (value) {
      if (value.name !== 'personal') {
        this.showNav = true
      } else {
        this.showNav = false
      }
    }
  }
}
</script>

<style lang="scss" scoped>
    .layout {
        height: 100vh;
        background-color: $-light;
        overflow: hidden;
        display: flex;
        nav {
            display: none;
            .logo{
                position: absolute;
                left: 50%;
                transform: translateX(-50%);
            }
        }

        aside {
            text-align: center;
            position: relative;
            width: 150px;
            height: 100%;
            background-color: $-theme;
            padding: $-padding;
            background-image: url("../../assets/img/spllit.png");
            .avatar {
                width: 120px;
                height: 120px;
                margin: 50px auto;
                border-radius: 50%;
                box-shadow: $-box-shadow;
            }

            ul {
                margin: 30px 0;
            }

            li {
                height: 50px;
                line-height: 50px;
                font-size: 14px;
                color: #eee;
                text-align: center;
                cursor: pointer;

                &:hover {
                    color: #fff;
                }
            }
        }

        .submenu {
            width: 120px;
            height: 100%;
            background-color: mix(#000, $-theme, 10%);
            background-image: url("../../assets/img/spllit.png");
            box-shadow: inset 5px 0 5px rgba(0, 0, 0, 0.1);
        }

        .menu-btn{
            float: left;
            line-height: 60px;
            width: 60px;
            font-size: 24px;
            color: #f5f5f5;
        }
    }

    .menu-list {
        margin: 30px 0;

        li {
            height: 50px;
            line-height: 50px;
            font-size: 14px;
            color: #eee;
            text-align: center;
            cursor: pointer;

            &:hover {
                color: #fff;
            }
        }
    }

    .content {
        flex: 1;
        overflow: hidden;
        height: 100%;

        & > div {
            height: 100%;
        }
    }

    @media (max-width: 780px) {
        .layout{
            display: block;
            aside {
                display: none;
            }
            .submenu {
                display: none;
            }
            nav{
                display: block;
                width: 100%;
                height: 60px;
                position: fixed;
                background-color: $-theme;
                z-index: 10;
                transition: all .3s linear;
                background-image: url("../../assets/img/spllit.png");
                box-shadow: 0 0 10px rgba(0,0,0,.3);
                overflow: hidden;
                opacity: 0;
                text-align: center;
                .logo{
                   padding: $-padding;
                }
                &.show{
                    opacity: 1;
                }
            }
        }
    }
</style>
