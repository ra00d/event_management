import { create } from "zustand";

export const authStore = create((_set, _get) => ({
	user: {
		name: "Guest",
		role: 1,
		token: "",
	},
}));
