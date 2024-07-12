import apiClient from "@/common/api";
import { Button } from "@/components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { AddEventT, eventSchema } from "@/lib/types/events";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
	Calendar,
	DollarSign,
	Ticket,
	Upload,
	User,
	XCircle,
} from "lucide-react";
import { useEffect, useState } from "react";
import { useFieldArray, useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import { toast } from "sonner";

const useEventState = () => {
	const { id } = useParams();
	const { data } = useQuery({
		queryKey: ["event" + id],
		queryFn: async () => {
			const res = await apiClient.get(`events/${id}`);
			return res.data;
		},
		enabled: id !== "new",
	});
	const { mutate, isPending } = useMutation({
		mutationFn: (data: AddEventT) => {
			if (id === "new") return apiClient.postForm("/events", data);
			return apiClient.putForm(`/events/${id}`, data);
		},
		onSuccess: (res) => {
			const data = res.data;
			toast.success(`event ${id === "new" ? "added" : "updated"}  successfuly`);
			setEventImages(data?.images || []);
			form.reset({
				...data,
				date: new Date(data.date),
				images: [],
				category: "default",
			});
		},
		onError() {
			toast.error("sorry somwthing went wrong");
		},
	});
	const [eventImages, setEventImages] = useState<string[]>([]);
	const form = useForm<AddEventT>({
		resolver: zodResolver(eventSchema),
		defaultValues: {
			date: new Date(),
			images: [],
			// deleted_images: [""],
			category: "default",
		},
	});
	const handelSubmit = async (values: AddEventT) => {
		// values.ticket_price = Number(values.ticket_price);
		// values.tickets = Number(values.tickets);
		mutate(values);
	};
	const images = useFieldArray({
		control: form.control, // control props comes from useForm (optional: if you are using FormProvider)
		name: "images", // unique name for your Field Array
	});

	const deletedImages = useFieldArray({
		control: form.control, // control props comes from useForm (optional: if you are using FormProvider)
		//@ts-ignore
		name: "deleted_images", // unique name for your Field Array
	});
	useEffect(() => {
		form.reset({
			...data,
			date: new Date(data?.date),
			images: [],
			deleted_images: [],
			category: "default",
		});
		if (data?.images) setEventImages(data?.images);
	}, [data]);
	useEffect(() => {
		console.log(form.formState.errors);
	}, [form.formState.errors]);
	return {
		form,
		isPending,
		handelSubmit,
		fields: images,
		deletedImages,
		images: eventImages,
		setEventImages,
	};
};
export const NewEvent = () => {
	const {
		form,
		isPending,
		handelSubmit,
		fields,
		deletedImages,
		images,
		setEventImages,
	} = useEventState();
	return (
		<div className="container space-y-5 mx-auto bg-card border  shadow-md rounded-md p-6">
			<h1>New Event</h1>
			<Form {...form}>
				<form
					className="grid grid-cols-1 md:grid-cols-2 gap-5"
					onSubmit={form.handleSubmit(handelSubmit)}
				>
					<FormField
						name="name"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem>
									<FormLabel>Name</FormLabel>
									<FormControl>
										<Input {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>

					<FormField
						name="organizer"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem className="">
									<FormLabel className="flex gap-2 items-center">
										Organizer
										<User />
									</FormLabel>
									<FormControl>
										<Input {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>

					<FormField
						name="address"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem className="">
									<FormLabel>Address</FormLabel>
									<FormControl>
										<Input {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>
					<FormField
						name="date"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem>
									<FormLabel className="flex items-center gap-2">
										Date <Calendar />
									</FormLabel>
									<FormControl>
										<Input
											{...field}
											type="datetime-local"
											// value={field.value?.toString()}
											value={form.watch("date")?.toString()}
											onChange={(val) => {
												field.onChange(new Date(val.target.value));
											}}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>
					<FormField
						name="description"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem className="col-span-2">
									<FormLabel>Description</FormLabel>
									<FormControl>
										<Textarea {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>
					<FormField
						name="tickets"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem>
									<FormLabel className="flex items-center gap-2">
										Tickets <Ticket />
									</FormLabel>
									<FormControl>
										<Input type="number" {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>

					<FormField
						name="ticket_price"
						control={form.control}
						render={({ field }) => {
							return (
								<FormItem>
									<FormLabel className="flex items-center gap-2">
										Ticket Price
										<DollarSign />
									</FormLabel>
									<FormControl>
										<Input type="number" {...field} />
									</FormControl>
									<FormMessage />
								</FormItem>
							);
						}}
					/>
					<div className="col-span-2 flex gap-5">
						<input
							type="file"
							accept="image/*"
							hidden
							id="advertisement-img"
							multiple
							onChange={(e) => {
								const files = e.target.files;
								if (files) {
									const arr = Array.from(files);
									arr.forEach((file) => {
										fields.append(file);
									});
								}
								e.target.files = null;
							}}
						/>
						<div className="w-[128px] h-[128px] relative flex justify-center items-center bg-background rounded-md shadow-md">
							<Label htmlFor="advertisement-img" className="cursor-pointer">
								<Upload />
							</Label>
						</div>
						{images.map((img: string) => (
							<div
								key={img}
								className="flex flex-col gap-2 justify-center items-center"
							>
								<div className="w-[128px] h-[128px] relative flex justify-center items-center bg-background rounded-md shadow-md overflow-clip ">
									<img
										src={"http://localhost:8080" + img}
										className="h-auto w-full object-fill"
										alt="imag"
									/>
								</div>
								<Button
									type="button"
									// asChild
									variant="ghost"
									size="icon"
									onClick={() => {
										// const deletedImg = fields
										// 	?.filter((i) => i.imgUrl === img)
										// 	.map((i) => i.id);
										// if (deletedImg && deletedImg?.length > 0)
										// 	deletedImages.append(deletedImg[0]);
										//
										setEventImages((prev) => prev.filter((i) => i != img));
										deletedImages.append(img);
									}}
								>
									<XCircle fill="red" />
								</Button>
							</div>
						))}
						<>
							{form.watch("images")?.map((img, idx) => {
								return (
									<div
										key={img.id}
										className="flex flex-col gap-2 justify-center items-center"
									>
										<div className="w-[128px] h-[128px] relative flex justify-center items-center bg-background rounded-md shadow-md overflow-clip ">
											<img
												src={URL.createObjectURL(img)}
												className="h-auto w-full object-fill"
												alt="imag"
											/>
										</div>
										<Button
											type="button"
											// asChild
											variant="ghost"
											size="icon"
											onClick={() => {
												// const deletedImg = fields
												// 	?.filter((i) => i.imgUrl === img)
												// 	.map((i) => i.id);
												// if (deletedImg && deletedImg?.length > 0)
												// 	deletedImages.append(deletedImg[0]);
												fields.remove(idx);
											}}
										>
											<XCircle fill="red" />
										</Button>
									</div>
								);
							})}
						</>{" "}
					</div>
					<Button loading={isPending}>save</Button>
				</form>
			</Form>
		</div>
	);
};
