import apiClient from "@/common/api";
// biome-ignore lint/style/useImportType: <explanation>
import { TEvent } from "@/lib/types/events";

export const getEvents = async () => {
	const res = await apiClient.get("/events");
	return res.data as TEvent[];
};
