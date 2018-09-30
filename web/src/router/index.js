import Vue from 'vue'
import Router from 'vue-router'
import layout from '@/components/Layout'
const _import = require('./_import_' + process.env.NODE_ENV)

Vue.use(Router)
 
export default new Router({
  // mode: 'history',
  // base: __dirname,
  // linkActiveClass: 'active', // 更改激活状态的Class值
  routes: [
    {
      path: '/',
      name: 'index',
      component: layout,
      children: [
        {
          path: 'test',
          name: 'test',
          component: _import('test/index')
        }
      ]
    }
  ]
})
