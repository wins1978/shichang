const MyPlugin = {
  // install方法是必需的
  // 包含两个参数：Vue 构造器，一个可选的选项对象
  install (Vue, options) {
    // 使用minxin将功能注入所有组件
    Vue.mixin({
      // 添加到mixin中的任何内容将被注入到所有组件中。
      // 在这个例子中， mounted() 方法将在组件被挂载到DOM后调用
      mounted () {
        console.log('Mounted!')
      }
    })
  }
}

export default MyPlugin
