<template>
  <div class="me">
    <div class="tabs">
      <ul>
        <!--<router-link :to="{name: 'Default'}" tag="li" exact>默认内容</router-link>-->
        <router-link :to="{name: 'Collection'}" tag="li" >我的收藏</router-link>
        <router-link :to="{name: 'Trace'}" tag="li">我的足迹</router-link>
      </ul>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>
<script type="text/ecmascript-6">
 
</script>
