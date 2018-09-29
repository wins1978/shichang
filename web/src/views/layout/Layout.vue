<template>
<div class="sidebar-container">
  <LeftMenu></LeftMenu>
  <div class="main-container"></div>
</div>
</template>

<script>
import {
  default as LeftMenu
} from '../../components/LeftMenu'

export default {
  components: {
    LeftMenu
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  @import "../../styles/layout.scss";
</style>

