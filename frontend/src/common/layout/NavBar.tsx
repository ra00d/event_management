import { Button } from "@/components/ui/button";
import { MobileMenu } from "./MobileMenu";
import { Link } from "react-router-dom";

const NavBar = () => {
	return (
		<div className=" px-10 text-primary py-2 sticky w-full top-0">
			<MobileMenu />
			<nav className="hidden md:block">
				<div className=" mx-auto flex justify-between items-center">
					<div className=" font-bold text-lg">
						<strong className="text-red-500 text-3xl">M</strong>Event
					</div>
					<div className="space-x-4">
						<Link to="/events" className="">
							Events
						</Link>
					</div>
					<Button className="">
						<a href="/login" className="">
							Login
						</a>
					</Button>
				</div>
			</nav>
		</div>
	);
};
export default NavBar;
