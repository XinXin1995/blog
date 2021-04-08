import Vue from 'vue'

import element from 'element-ui'
import vuescroll from 'vuescroll'
import hljs from 'highlight.js'
import javascript from 'highlight.js/lib/languages/javascript'
import './filters'
import lazyload from 'vue-lazyload'

import './assets/reset.scss'
import 'element-ui/lib/theme-chalk/index.css'
import 'highlight.js/styles/atom-one-light.css'
import '@/assets/style/iconfont.css'

hljs.registerLanguage('javascript', javascript)
Vue.use(element)
Vue.use(vuescroll)
Vue.use(lazyload, {
  preLoad: 1,
  error: require('@/assets/img/loadingerror.svg'),
  loading: require('@/assets/img/rings.svg'),
  attempt: 1
})
