const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: [
    './root.js'
  ],
  output: {
    path: path.resolve(__dirname, 'dest'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['env', 'react']
          }
        }
      },
      {
        test: /\.css$/,
        use: [
          'style-loader',
          {loader: 'css-loader', options: {url: false}},
        ],
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'GraphQL RealTimeChat',
      template: 'root.html'
    })
  ],
  devServer: {
    contentBase: path.join(__dirname, 'dest'),
    port: 9000
  },
  performance: { hints: false }
};
