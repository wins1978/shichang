<template>
  <div id="app">
    <v-app id="inspire">
      <v-navigation-drawer app permanent>
        <v-toolbar flat>
          <v-list>
            <v-list-tile>
              <v-list-tile-title class="title">
                石场报表
              </v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-toolbar>

        <v-divider></v-divider>

        <v-list dense class="pt-0">
           <v-list-tile @click="redirect('home')">
          <v-list-tile-action>
            <v-icon>home</v-icon>
          </v-list-tile-action>
          <v-list-tile-title>生成报表</v-list-tile-title>
        </v-list-tile>
          <v-list-group
            prepend-icon="account_circle"
            value="true"
            clipped="true"
          >
            <v-list-tile slot="activator">
              <v-list-tile-title>数据下载</v-list-tile-title>
            </v-list-tile>
            <v-list-tile
              v-for="(admin, i) in admins"
              :key="i"
            >
              <v-list-tile-title v-text="admin[0]"></v-list-tile-title>
              <v-list-tile-action>
                <v-icon v-text="admin[1]"></v-icon>
              </v-list-tile-action>
            </v-list-tile>
          </v-list-group>
          <v-list-tile v-for="item in items" :key="item.title" @click="redirect(item.page)">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ item.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>
    </v-app>
  </div>
</template>

<script>
export default {
  name: 'LeftNav',
  data () {
    return {
      admins: [
        ['微信日报', 'people_outline'],
        ['客户日报', 'people_outline'],
        ['周报表', 'people_outline'],
        ['月出库核算表', 'people_outline'],
        ['自定义报表', 'people_outline']
      ],
      items: [
        { title: '客户资料', page: 'Customer', icon: 'dashboard' },
        { title: '石粉型号', page: 'Stone', icon: 'dashboard' },
        { title: '客户价格维护', page: 'Price', icon: 'dashboard' },
        { title: '客户预收款', page: 'Deposit', icon: 'dashboard' },
        { title: '历史数据更新', page: 'FixData', icon: 'dashboard' }
      ],
      right: null
    }
  },
  methods: {
    redirect (item) {
      var tag = '/' + item
      this.$router.push({path: tag})
    }
  }
}
</script>
