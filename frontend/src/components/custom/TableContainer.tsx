import { ReactNode } from "react";

export const TableContainer = (props: { children: ReactNode }) => {
	return (
		<div className="bg-card border rounded-md min-h-[calc(100vh-20vh)] shadow-md overflow-clip">
			{props.children}
		</div>
	);
};
