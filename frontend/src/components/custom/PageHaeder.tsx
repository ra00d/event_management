import { Link } from "react-router-dom";
import { Button } from "../ui/button";
import { Plus } from "lucide-react";
export type PageHaederProps = { title: string; newPage?: string };
export const PageHaeder = ({ title, newPage }: PageHaederProps) => {
	return (
		<div className="flex my-3 justify-between items-center">
			<h1>{title}</h1>
			<div>
				<Button variant={"ghost"}>
					<Link to={newPage || ""}>
						<Plus />
					</Link>
				</Button>
			</div>
		</div>
	);
};
