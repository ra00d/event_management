import { z } from "zod";

const eventData = {
	id: 1,
	name: "event name",
	date: "2023-11-24 23:21:43",
	address: "Taiz",
	category: "sports",
	description: "this is the description",
	images: [""],
	number_of_tickets: 400,
	avialable_tickets: 400,
	organizer_id: "Organizer123",
	lat: 40.7128,
	lang: -74.006,
};
export type TEvent = typeof eventData;
export const eventSchema = z.object({
	name: z.string(),
	description: z.string(),
	category: z.string(),
	date: z.coerce.date(),
	address: z.string(),
	tickets: z.coerce.number(),
	ticket_price: z.coerce.number(),
	organizer: z.string(),
	images: z.array(z.any().or(z.string())).optional(),
	deleted_images: z.array(z.string()).optional(),
});
export type AddEventT = z.infer<typeof eventSchema>;
