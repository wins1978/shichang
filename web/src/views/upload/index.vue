<template>
  <div>
    <h1>{{ msg }}</h1>
    <upload-file @changed-files="uploadAction($event)"></upload-file>
  </div>
</template>

<script>
import UploadFile from 'vue-simple-upload-component'

export default {
  name: 'test',
  data() {
    return {
      msg: 'upload',
      files: []
    }
  },
  components: {
    UploadFile
  },
  created() {
  },
  methods: {
    /**
     * Has changed
     * @param  Object|undefined   newFile   只读
     * @param  Object|undefined   oldFile   只读
     * @return undefined
     */
    inputFile(newFile, oldFile) {
      console.log('response', newFile.response)
      if (newFile && oldFile && !newFile.active && oldFile.active) {
        // 获得相应数据
        console.log('response', newFile.response)
        if (newFile.xhr) {
          //  获得响应状态码
          console.log('status', newFile.xhr.status)
        }
      }
    },
    /**
     * Pretreatment
     * @param  Object|undefined   newFile   读写
     * @param  Object|undefined   oldFile   只读
     * @param  Function           prevent   阻止回调
     * @return undefined
     */
    inputFilter(newFile, oldFile, prevent) {
      if (newFile && !oldFile) {
        // 过滤不是图片后缀的文件
        if (!/\.(xls|xlsx)$/i.test(newFile.name)) {
          return prevent()
        }
      }

      // 创建 blob 字段 用于图片预览
      newFile.blob = ''
      var URL = window.URL || window.webkitURL
      if (URL && URL.createObjectURL) {
        newFile.blob = URL.createObjectURL(newFile.file)
      }
    }
  }
}
</script>
