import { Outlet } from "react-router-dom";
import NavBar from "./NavBar";
const Layout = (_props: { pageName?: string }) => {
	// if (user === "USER") {
	// 	return <Navigate to={"/login"} />;
	// }
	return (
		<div className="flex flex-col h-screen">
			<NavBar />
			<div className="flex-1 px-10 h-full">
				<Outlet />
			</div>
		</div>
	);
};
export default Layout;
