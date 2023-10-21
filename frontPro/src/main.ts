import { createSSRApp } from 'vue'
import TnIcon from '@tuniao/tnui-vue3-uniapp/components/icon/src/icon.vue'
import App from './App.vue'
import request from './utils/requests/request'
import vcache from './utils/requests/vcache'
import vcommon from './utils/requests/vcommon'

import './static/style/index.scss'
import 'vant/es/toast/style';
import 'vant/es/dialog/style';
import 'vant/es/notify/style';

export function createApp() {
  const app = createSSRApp(App)
  app.config.globalProperties.$request = request;
	app.config.globalProperties.$vcache = vcache;
	app.config.globalProperties.$vcommon = vcommon;
  app.component('TnIcon', TnIcon)

  return {
    app,
  }
}
