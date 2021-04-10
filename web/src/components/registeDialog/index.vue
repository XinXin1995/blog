<template>
    <el-dialog title="注册" :visible.sync="vis" width="400px" :before-close="beforeClose" top="30px">
        <el-form ref="register" status-icon label-width="80px" :model="param" :rules="rules">
            <el-form-item label="用户名" prop="username">
                <el-input v-model="param.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input type="password" v-model="param.password"></el-input>
            </el-form-item>
            <el-form-item label="重复密码" prop="repeatPwd">
                <el-input type="password" v-model="param.repeatPwd"></el-input>
            </el-form-item>
            <el-form-item label="头像" prop="avatar">
                <upload :file-type="1" :file="param.avatar" @success="handleUploadSuccess"
                        @remove="handleUploadRemove"/>
            </el-form-item>
        </el-form>
        <div slot="footer">
            <el-button type="default" @click="handleClose">取消</el-button>
            <el-button type="primary" @click="handleConfirm">确定</el-button>
        </div>
    </el-dialog>
</template>

<script>
import Upload from '@/components/upload/image'
import { AddUser } from '@/api/user'
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
    const validName = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入用户名'))
      } else if (value.length < 3) {
        callback(new Error('用户名至少3个字符'))
      } else {
        callback()
      }
    }
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else if (value.length < 6) {
        callback(new Error('密码长度至少6位'))
      } else {
        if (this.param.repeatPwd !== '') {
          this.$refs.register.validateField('repeatPwd')
        }
        callback()
      }
    }
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.param.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      rules: {
        username: [
          {
            validator: validName,
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: validatePass,
            trigger: 'blur'
          }
        ],
        repeatPwd: [
          {
            validator: validatePass2,
            trigger: 'blur'
          }
        ],
        avatar: []
      },
      param: {
        username: '',
        password: '',
        repeatPwd: '',
        avatar: ''
      },
      baseURL: 'https://store.wuchangxin.club/user/1607062845_avatar.jpg'
    }
  },
  methods: {
    ...mapActions(['login']),
    beforeClose (done) {
      this.param = {
        username: '',
        password: '',
        repeatPwd: ''
      }
      this.$refs.register.resetFields()
      this.$emit('update:visible', false)
      done()
    },
    handleClose () {
      this.$emit('update:visible')
    },
    handleUploadSuccess (data) {
      console.log(data)
      this.param.avatar = data.url
    },
    handleUploadRemove () {
      this.param.avatar = ''
    },
    handleConfirm () {
      this.$refs.register.validate(valid => {
        if (valid) {
          this.param.avatar = this.param.avatar || this.baseURL
          AddUser(this.param).then(res => {
            if (res.code === 0) {
              this.login(this.param)
              this.handleClose()
            }
          })
        }
      })
    }
  },
  components: {
    Upload
  }
}
</script>

<style lang="scss" scoped>
</style>
