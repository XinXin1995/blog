<template>
    <div>
        <div class="login-register" v-if="!(user && user.username)">
            <div class="login" @click="loginVisible = true">
                登录
            </div>
            <div class="register" @click="registerVisible = true">
                注册
            </div>
        </div>
        <div class="login-register" v-else>
            <div>
                {{user.username}}
            </div>
        </div>
        <login-dialog :visible.sync="loginVisible" />
        <register-dialog :visible.sync="registerVisible" />
    </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import { getToken } from '@/utils/auth'
import LoginDialog from '@/components/loginDialog'
import registerDialog from '@/components/registeDialog'
export default {
  computed: {
    ...mapGetters(['user'])
  },
  data () {
    return {
      loginVisible: false,
      registerVisible: false
    }
  },
  methods: {
    ...mapActions(['login', 'getUser']),
    init () {
      if (getToken()) {
        this.getUser()
      }
    }
  },
  created () {
    this.init()
  },
  components: {
    LoginDialog,
    registerDialog
  }
}
</script>

<style lang="scss" scoped>
    .login-register{
        width: 100%;
        height: 40px;
        line-height: 40px;
        display: flex;
        position: absolute;
        left: 0;
        bottom: 0;
        color: $-bgc;
        & > div {
            height: 100%;
            flex: 1;
            text-align: center;
            cursor: pointer;
            &:hover{
                color: #fff;
            }
        }
        .register{
            position: relative;
            &:after{
                content: '';
                display: block;
                width: 1px;
                height: 14px;
                position: absolute;
                left: -0.5px;
                background-color: $-bgc;
                top: 50%;
                margin-top: -7px;
            }
        }
    }
</style>
