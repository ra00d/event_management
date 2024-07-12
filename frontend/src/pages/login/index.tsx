import { Heading1 } from "@/components/ui/Headings";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
// import { Label } from "@/components/ui/label";
import { Apple, Facebook, Loader2, Twitter } from "lucide-react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Link, useNavigate } from "react-router-dom";
import * as z from "zod";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { type ILogin, login } from "./fetchers";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import type { AxiosError } from "axios";

const formSchema = z.object({
	email: z
		.string()
		.min(10, { message: "to short email" })
		.email({ message: "This is not a valid email" }),
	password: z.string().min(6, { message: "Too short password" }),
});

export default function LoginPage() {
	const navigate = useNavigate();
	const [loginError, setLoginError] = useState("");
	const { mutate, isPending } = useMutation({
		mutationFn: (data: ILogin) => login(data),
		onError: (error: AxiosError<{ message: string }>) => {
			setLoginError(error.response?.data?.message || "");
		},
		onSuccess: () => {
			navigate("/", {
				replace: true,
			});
		},
	});
	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	});
	const onSubmit = (values: ILogin) => {
		mutate(values);
	};
	return (
		<div className="container flex items-center justify-center h-screen p-10 ">
			<div className="grid md:grid-cols-2 w-full md:h-full my-10 overflow-clip rounded">
				<div className="hidden md:flex justify-center items-center bg-gray-200">
					<Heading1>MEvent</Heading1>
				</div>
				<div className="flex flex-col gap-7 py-5 md:py-0  justify-center items-center bg-gray-100 ">
					<Heading1>Login</Heading1>
					<div className="flex  flex-col gap-4 w-full px-5 md:px-32">
						<div className="flex gap-1 justify-center">
							<Facebook />
							<Twitter />
							<Apple />
						</div>
						{loginError.length > 0 && (
							<div className="bg-red-500 rounded-md px-4 py-2 text-start text-white">
								<span className="text-xl  ">{loginError}</span>
							</div>
						)}

						<Form {...form}>
							<form
								onSubmit={form.handleSubmit(onSubmit)}
								className="flex flex-col gap-4 "
							>
								<FormField
									control={form.control}
									name="email"
									render={({ field }) => (
										<FormItem>
											<FormLabel>Username</FormLabel>
											<FormControl>
												<Input {...field} placeholder="enter your email" />
											</FormControl>
											{/*<FormDescription>
													This is your email address
												</FormDescription>*/}
											<FormMessage />
										</FormItem>
									)}
								/>
								<div>
									<FormField
										control={form.control}
										name="password"
										render={({ field }) => (
											<FormItem>
												<FormLabel>Password</FormLabel>
												<FormControl>
													<Input
														type="password"
														{...field}
														placeholder="enter your password"
													/>
												</FormControl>
												{/*<FormDescription>
													This is your email address
												</FormDescription>*/}
												<FormMessage />
											</FormItem>
										)}
									/>{" "}
									<div className="w-full flex justify-end">
										<Link to={"#"} className="text-blue-500 text-end">
											Forget password?
										</Link>
									</div>
								</div>
								<div className="mt-3">
									<Button className="w-full" disabled={isPending}>
										{isPending ? <Loader2 /> : "Login"}
									</Button>
								</div>
							</form>
						</Form>
					</div>
				</div>
			</div>
		</div>
	);
}
