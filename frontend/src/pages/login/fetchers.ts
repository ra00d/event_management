import apiClient from "@/common/api";

export interface ILogin {
	email: string;
	password: string;
	remember_me?: boolean;
}
export const login = async (data: ILogin) => {
	// await apiClient.get("/csrf").then((res)=>{
	//
	// });
	await apiClient.post("/auth/log-in", data);
};
