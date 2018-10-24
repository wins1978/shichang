<template>
  <div>
    <h1>upload</h1>
    <uploader :options="options" class="uploader-example">
    <uploader-unsupport></uploader-unsupport>
    <uploader-drop>
      <p>Drop files here to upload or</p>
      <uploader-btn :attrs="attrs">select files</uploader-btn>
    </uploader-drop>
    <uploader-list></uploader-list>
  </uploader>
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
        options: {
          target: 'http://localhost:8080/Upload/UpFile',
          testChunks: false
        },
        attrs: {
          accept: 'document/*'
        }
      }
  },
  components: {
  },
  created() {
  },
  methods: {
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
