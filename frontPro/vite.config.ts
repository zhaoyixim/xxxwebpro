import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'

import Components from 'unplugin-vue-components/vite';
import { VantResolver } from 'unplugin-vue-components/resolvers';
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [uni(),Components({
	      resolvers: [VantResolver()],
	    }),],
			 //  server: {
			 //    //host:'0.0.0.0',
				// proxy: {
				//       "/api": {
				//         target: "http://127.0.0.1:8088",
				//         changeOrigin: true,
				//         rewrite: (path) => path.replace(/^\/api/, ''), 
				//       },
				//     },
			 //  },
  base: './',
})
