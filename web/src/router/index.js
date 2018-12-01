import Vue from 'vue'
import Router from 'vue-router'
import Upload from '@/components/Upload'
import Customer from '@/components/Customer'
import Vuetify from 'vuetify'

Vue.use(Vuetify)
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Upload',
      component: Upload
    },
    {
      path: '/Customer',
      name: 'Customer',
      component: Customer
    }
  ]
})
