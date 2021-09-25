const path = require('path');
const webpack = require('webpack');

const Dotenv = require('dotenv-webpack');

module.exports = {
  mode: 'development',
  plugins: [new Dotenv()],
  entry: {
    "firebase-messaging-sw": './src/service-worker.ts',
    "script": './src/script.ts'
  },
  devtool: 'inline-source-map',
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'dist'),
  }
};