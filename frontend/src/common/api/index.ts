import axios from "axios";

const getCsrfToken = () => {
	let token;
	token = localStorage.getItem("csrf_token");
	if (token == null || token === undefined) {
		axios.get("http://localhost:8080/csrf").then((res) => {
			console.log(res.data);
			token = res.data?.csrf;
			localStorage.setItem("csrf_token", token);
			// csrfToken = token;
		});
	}
	return token;
};
axios.defaults.headers.common.Accept = "application/json";
axios.defaults.withCredentials = true;
axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

const apiClient = axios.create({
	baseURL: "http://localhost:8080",
	headers: {
		"Content-Type": "application/json",
		Accept: "application/json",
		"X-Csrf-Token": getCsrfToken(),
	},
});

export default apiClient;
