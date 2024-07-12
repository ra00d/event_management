import * as React from "react";

import { cn } from "@/lib/utils";

export interface InputProps
	extends React.InputHTMLAttributes<HTMLInputElement> {}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
	({ className, type, ...props }, ref) => {
		return (
			<input
				type={type}
				className={cn(
					`flex h-10 w-full rounded-md border border-input bg-background
          px-3 py-2 text-sm ring-offset-background file:border-0 
          file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground
          focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring 
          focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50`,
					className,
				)}
				ref={ref}
				{...props}
			/>
		);
	},
);
Input.displayName = "Input";
export type InputWithIconProps = InputProps & {
	icon: React.ReactNode;
	inputClassName?: string;
};
const IconedInput = React.forwardRef<HTMLInputElement, InputWithIconProps>(
	({ className, inputClassName, icon, type, ...props }, ref) => {
		return (
			<div
				className={cn(
					"flex bg-background border border-input rounded-md w-full items-center overflow-hidden",
					"focus-within:outline-none focus-within:ring-2 focus-within:ring-ring focus-within:ring-offset-2",
					className,
				)}
			>
				<Input
					ref={ref}
					type={type}
					className={cn(
						"bg-transparent pointer-events-auto border-none focus-visible:ring-0 focus-visible:ring-offset-0",
						inputClassName ?? "",
					)}
					{...props}
				/>
				{icon}
			</div>
		);
	},
);
export { Input, IconedInput };
