'use strict';

var pathto = function (file) {
    return ('./public/' + file);
};
var scripts = {
    src: pathto('javascripts/*.js'),
    dest: pathto('dist')
};
var styles = {
    src: pathto('stylesheets/[name].css'),
    // scss: pathto('stylesheets/[name]/[name].scss'),
    dest: pathto('dist')
};

var path = require('path');
var glob = require('glob');
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const extractSass = new ExtractTextPlugin({
    filename: '[name].css', // 把路径重新定义到dist目录下，而不是css里面
});

module.exports = {
    entry: glob.sync(scripts.src), //入口文件,从项目根目录指定
    output: { //输出路径和文件名，使用path模块resolve方法将输出路径解析为绝对路径
        path: path.resolve(__dirname, scripts.dest), //将js文件打包到dist/js的目录
        filename: "app.js"
    },
    module: { // 如何处理项目中不同类型的模块
        rules: [ // 用于规定在不同模块被创建时如何处理模块的规则数组
            {
                test: /\.scss$/,
                use: extractSass.extract({
                    use: [{
                        loader: "css-loader"
                    }, {
                        loader: "sass-loader"
                    }],
                    fallback: "style-loader"
                })
            },
        ]
    },
    plugins: [
        extractSass,
        new CleanWebpackPlugin()
    ]
};