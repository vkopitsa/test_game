const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
//const glob = require("glob");
const ClosurePlugin = require('closure-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');

module.exports = (env, argv) => {

    var mode = process.env.NODE_ENV || argv.mode;

    return {
        entry: {
            index: './src/app.ts',

            // closure: "./closure-library/closure/goog/base.js",

            protobuf: [
                "./vender/protobuf/js/binary/constants.js",
                "./vender/protobuf/js/binary/utils.js",
                "./vender/protobuf/js/binary/arith.js",
                "./vender/protobuf/js/binary/encoder.js",
                "./vender/protobuf/js/binary/decoder.js",
                "./vender/protobuf/js/map.js",
                "./vender/protobuf/js/binary/writer.js",
                "./vender/protobuf/js/binary/reader.js",
                "./vender/protobuf/js/message.js",
            ],
        },
        mode: 'development',
        module: {
            rules: [{
                    test: /\.tsx?$/,
                    use: 'ts-loader',
                    exclude: /node_modules/
                },
                {
                    test: /\.css$/i,
                    use: ['style-loader', 'css-loader'],
                },
            ],
        },
        resolve: {
            extensions: ['.tsx', '.ts', '.js']
        },
        devtool: 'inline-source-map',
        plugins: [
            new CleanWebpackPlugin(),
            new HtmlWebpackPlugin({
                template: 'src/index.html',
                minify: {
                    collapseWhitespace: true,
                    removeComments: true
                },
            }),
            new CopyWebpackPlugin(
                [{
                    from: 'src/favicon.ico',
                    to: path.resolve(__dirname, 'dist')
                }, ]
            ),
            new ClosurePlugin.LibraryPlugin({
                closureLibraryBase: require.resolve(
                    './vender/closure-library/closure/goog/base'
                ),
                deps: [
                    require.resolve('./vender/closure-library/closure/goog/deps'),
                ],
            })
        ],
        output: {
            filename: '[name].[hash:8].js',
            chunkFilename: '[name].[hash:8].js',
            path: path.resolve(__dirname, 'dist')
        },
        stats: {
            warnings: false
        },
        optimization: {
            minimize: mode === 'production',
            minimizer: [
                new TerserPlugin({
                    cache: true,
                    parallel: true,
                    sourceMap: true,
                    extractComments: false,
                }),
            ],
            splitChunks: {
                chunks: 'all',
                minChunks: 2,
                minSize: 0,
            }
        },
        node: {
            fs: "empty",
            readline: "empty",
            gulp: "empty",
            jspb: "empty",
            goog: "empty",
        },
        devServer: {
            port: 9000,
            contentBase: './dist',
            proxy: {
                '/ws': {
                    target: 'ws://127.0.0.1:8080',
                    secure: false,
                    ws: true,
                },
            },
        },
    };;
};