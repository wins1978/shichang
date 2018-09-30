<template>
<div >
  <el-row >
    <el-col :span="44">
      <div class="sidebar-container">
         <LeftMenu></LeftMenu>
      <div class="main-container"></div>
      </div>
     
    </el-col>
    <el-col :span="44" class="full_height">
      <router-view/>
    </el-col>
  </el-row>
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

