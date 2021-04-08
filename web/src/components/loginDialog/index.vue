<template>
    <el-dialog :visible.sync="vis" :before-close="loginBeforeClose" width="400px" title="登录">
        <el-form ref="login" :rule="loginRule" :model="loginParam">
            <el-form-item label="用户名" prop="username">
                <el-input v-model="loginParam.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input type="password" v-model="loginParam.password" @keyup.native.enter="handleLogin"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer">
            <el-button type="default" @click="handleClose">取消</el-button>
            <el-button type="primary" :loading="loginLoading" @click="handleLogin">确定</el-button>
        </div>
    </el-dialog>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  props: {
    visible: Boolean
  },
  computed: {
    vis: {
      get () {
        return this.visible
      },
      set () {

      }
    }
  },
  data () {
    return {
      loginLoading: false,
      loginParam: {
        username: '',
        password: ''
      },
      loginRule: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    ...mapActions(['login']),
    loginBeforeClose (done) {
      this.loginParam = {
        username: '',
        password: ''
      }
      this.$emit('update:visible', false)
      done()
    },
    handleClose () {
      this.$emit('update:visible', false)
    },
    handleLogin () {
      this.loginLoading = true
      this.$refs.login.validate(valid => {
        if (valid) {
          this.login(this.loginParam).then(res => {
            if (res.code === 0) {
              this.isAuthed = true
              this.handleClose()
            }
            this.loginLoading = false
          }).catch(() => {
            this.loginLoading = true
          })
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
