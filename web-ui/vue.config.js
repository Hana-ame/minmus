const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api/': {
        target: 'http://127.0.8.1:8080/'
      },
    //   "/userapi": {
    //     target: 'http://localhost:3080/',
    //     pathRewrite: {'^/userapi' : '/api'}
    //   }
    }
  }

})
