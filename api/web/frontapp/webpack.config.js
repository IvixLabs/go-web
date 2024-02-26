const path = require('path');

module.exports = {
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      }
    ]
  },
  entry: './index.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, '../../../internal/controller/http/resources'),
  },
};