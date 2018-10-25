<template>
  <div>
    <h1>upload</h1>
    <el-upload  :action="UploadUrl()"  :on-success="UploadSuccess">
      <el-button size="small" type="primary" >点击上传</el-button>
    </el-upload>
  </div>
</template>

<script>
import Vue from 'vue'
import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default {
  name: 'test',
  data() {
    return {
    }
  },
  components: {
  },
  created() {
  },
  methods: {
    UploadUrl() {
      return 'http://localhost:8080/Upload/UpFile'
    },
    UploadSuccess() {
    }
  }
}
</script>

<style>
  .uploader-example {
    width: 880px;
    padding: 15px;
    margin: 40px auto 0;
    font-size: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, .4);
  }
  .uploader-example .uploader-btn {
    margin-right: 4px;
  }
  .uploader-example .uploader-list {
    max-height: 440px;
    overflow: auto;
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
