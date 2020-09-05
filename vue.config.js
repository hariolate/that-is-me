const path = require("path");

module.exports = {
    outputDir: path.resolve(__dirname, "./static"),
    publicPath: "./",
    configureWebpack: {
        resolve: {
            alias: {
                vue$: 'vue/dist/vue.esm.js'
            }
        }
    }
};