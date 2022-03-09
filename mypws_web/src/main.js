import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false

import "bootswatch/dist/litera/bootstrap.min.css";

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')