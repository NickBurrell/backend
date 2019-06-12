var path = require('path')

module.exports = {
	context: __dirname,

	mode: "production",

	entry: {
		index: './src/js/index.jsx'
	},
	
	output: {
		path: path.resolve('./public/js'),
		filename: '[name].js'
	},
	devtool: "source-map",

	module: {
		rules: [
			{
				test: /\.jsx?$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader'
				}
			},
			{
				test: /\.tsx?$/,
				exclude: /node_modules/,
				use: {
					loader: 'ts-loader'
				}
			}
		]
	},
	resolve: {
		extensions: ['.ts', '.tsx', '.js']
	},
	externals: {
		"react": "React",
		"react-dom": "ReactDOM"
	}
}
