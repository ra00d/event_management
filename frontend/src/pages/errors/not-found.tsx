import { Link } from "react-router-dom";

export const NotfoundPage = () => {
	return (
		<div className="flex flex-col items-center justify-center">
			<h1 className="text-5xl font-bold">404</h1>
			<span>This page does not exist</span>
			<Link to={"/"}>Home</Link>
		</div>
	);
};
