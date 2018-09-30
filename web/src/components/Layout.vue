<template>
<div>
<div >
  <el-row>
    <el-col :span="44" class="sidebar-container">
      <LeftMenu></LeftMenu>
      <div class="main-container"></div>
    </el-col>
    <el-col :span="44">
      <router-view/>
    </el-col>
  </el-row>
</div>
</div>

</template>

<script>
import {
  default as LeftMenu
} from '@/components/LeftMenu'
import {
  default as AppMain
} from '@/components/AppMain'

export default {
  components: {
    LeftMenu, AppMain
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  @import "@/styles/layout.scss";
</style>

