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
    entry: {
        shared: "./shared.js",
        video: {
            import: "./video.js",
            dependOn: "shared"
        }
    },
    output: {
        asyncChunks: true,
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, '../resources'),
    },
};