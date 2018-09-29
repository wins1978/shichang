import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '../views/layout/Layout'
const _import = require('./_import_' + process.env.NODE_ENV)

Vue.use(Router)

export const constantRouterMap = [
  {
    path: '/',
    component: Layout,
    children: [{
      path: '',
      component: _import('index/index'),
      name: '首页'
    }]
  },
  {
    path: '/test',
    component: Layout,
    children: [{
      path: 'index',
      component: _import('test/index'),
      name: '测试'
    }
    ]
  }
]

constantRouterMap.push({ path: '*', redirect: '/404' })

export default new Router({
  // mode: 'history', // require service support
    scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
