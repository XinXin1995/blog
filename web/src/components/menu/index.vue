<template>
    <transition name="slide">
    <div class="menu-dialog" v-show="visible" @click="handleClose">
        <div class="menu-wrap" v-show="visible" @click.stop>
            <logo />
            <div class="avatar">
                <img :src="user.avatar" alt="">
            </div>
            <ul>
                <li v-for="(item, index) in menu" :key="index" @click="handleDirect(item)">
                    <i :class="item.icon"></i>
                    {{item.name}}
                </li>
            </ul>
        </div>
    </div>
    </transition>
</template>

<script>
import { mapGetters } from 'vuex'
import Logo from '@/components/logo'
export default {
  props: {
    visible: Boolean,
    menu: Array
  },
  computed: {
    ...mapGetters(['user'])
  },
  methods: {
    handleClose () {
      this.$emit('update:visible', false)
    },
    handleDirect (item) {
      if (this.$route.name !== item.router) {
        this.$router.push({ name: item.router })
        this.handleClose()
      }
    }
  },
  components: {
    Logo
  }
}
</script>

<style lang="scss" scoped>
    .menu-dialog{
        width: 100%;
        height: 100%;
        position: fixed;
        z-index: 12;
        background: rgba(0,0,0,.3);
        .menu-wrap{
            overflow: hidden;
            width: 300px;
            height: 100%;
            background-color: $-theme;
            background-image: url('../../assets/img/spllit.png');
            transition: all .3s;
        }
        .logo{
            display: block;
            text-align: center;
            margin: $-margin * 2;
        }
        .avatar{
            width: 120px;
            height: 120px;
            margin: 30px auto;
            overflow: hidden;
            border-radius: 50%;

            img{
                width: 100%;
                height: 100%;
            }

        }
        li{
            height: 50px;
            line-height: 50px;
            text-align: center;
            font-size: 18px;
            color: #fff;
        }
    }
</style>
