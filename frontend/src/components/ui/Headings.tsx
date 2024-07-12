import { ReactNode } from "react";

export const Heading1 = ({ children }: { children: ReactNode }) => {
	return (
		<h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
			{children}
		</h1>
	);
};
export const Heading2 = ({ children }: { children: ReactNode }) => {
	return (
		<h1 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0">
			{children}{" "}
		</h1>
	);
};
export const Heading3 = ({ children }: { children: ReactNode }) => {
	return (
		<h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">
			{children}{" "}
		</h1>
	);
};
export const Heading4 = ({ children }: { children: ReactNode }) => {
	return (
		<h1 className="scroll-m-20 text-xl font-semibold tracking-tight">
			{children}
		</h1>
	);
};
export const Heading5 = ({ children }: { children: ReactNode }) => {
	return (
		<h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
			{children}
		</h1>
	);
};
