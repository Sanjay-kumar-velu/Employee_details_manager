import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': 'http://172.31.6.78:3030' // Forward /api calls to backend
    }
  }
})

/*import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: 'http://a4fe8263910484b3b878495fea7b776e-398127749.ap-southeast2.elb.amazonaws.com:3030',
        changeOrigin: true,
      }
    },
    allowedHosts: ['.elb.amazonaws.com']
  }
})*/
