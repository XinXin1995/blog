<template>
    <div class="">
        <div class="img"  v-if="url" @clic.stop.prevent>
            <img :src="url" class="avatar" alt=""/>
            <div class="close">
                <i class="el-icon-delete" @click="handleRemove"></i>
            </div>
        </div>
        <el-upload
                v-else
                class="picture-upload"
                :action="param.action"
                :data="param.data"
                :on-success="handleUploadSuccess"
                :show-file-list="false"
        >

            <div class="handler">
                    <i class="el-icon-plus avatar-uploader-icon"></i>
            </div>
        </el-upload>
    </div>
</template>

<script>
import { BASE_URL } from '@/config'
export default {
  props: {
    fileType: {
      type: Number,
      required: true
    },
    file: String
  },
  computed: {
    url: {
      get () {
        return this.file
      },
      set () {}

    }
  },
  data () {
    return {
      param: {
        action: BASE_URL + 'api/file/upload',
        data: {
          type: this.fileType
        }
      }
    }
  },
  methods: {
    handleUploadSuccess (res) {
      if (res.code === 0) {
        this.url = res.data
        this.$emit('success', res.data)
      }
    },
    handleRemove () {
      this.url = ''
      this.$emit('remove')
    }
  }
}
</script>

<style lang="scss" scoped>
    .picture-upload{
        display: inline-block;
        overflow: hidden;
        width: 200px;
        height: 200px;
        border: 1px dashed #ddd;
        border-radius: 10px;
        background-color: #fff;
        &:hover{
            border-color: #bbbbbb;
        }
        .handler{
            width: 198px;
            height: 198px;
            font-size: 30px;
            line-height: 200px;
            text-align: center;
            cursor: pointer;
            color: #dddddd;
            &:hover{
                color: #bbbbbb;
            }
        }
    }
    .img{
        border: 1px dashed #ddd;
        display: inline-block;
        border-radius: 10px;
        overflow: hidden;
        width: 198px;
        height: 198px;
        position: relative;
        .close{
            width: 100%;
            height: 100%;
            position: absolute;
            top: 0;
            left: 0;
            background-color: rgba(0,0,0,.5);
            font-size: 20px;
            color: #fff;
            display: none;
            line-height: 198px;
            text-align: center;
        }
        i{
            cursor: pointer;
            &:hover{
                color: $-active-color;
            }
        }
        img{
            width: 100%;
            min-height: 100%;
        }
        &:hover{
            .close{
                display: block;
            }
        }
    }
</style>
