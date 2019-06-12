import React from "react";
import ReactDOM from "react-dom";
import { Hello } from "./components/Hello";

const Index = () => {
	return <div>Hello React!</div>;
}

ReactDOM.render(<Index />, document.getElementById("index"));
