const path = require('path');

module.exports = {
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
        ]
    },
    devtool: 'inline-source-map',
    resolve: {
        extensions: ['.tsx', '.ts', '.js', '.css'],
    },
    entry: {
        shared: "./shared.js",
        video: {
            import: "./video.js",
            dependOn: "shared"
        },
        dashboard: "./dashboard.tsx",
        web: "./web.ts"
    },
    output: {
        asyncChunks: true,
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, '../internal/controller/http/handlers/web/static/resources'),
    },
};