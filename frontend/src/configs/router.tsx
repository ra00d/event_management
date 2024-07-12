import Layout from "@/common/layout/Layout";
import { NotfoundPage } from "@/pages/errors/not-found";
import { Events } from "@/pages/events/Events";
import { NewEvent } from "@/pages/events/new/NewEvent";
import LoginPage from "@/pages/login";
import { createBrowserRouter } from "react-router-dom";

const router = createBrowserRouter([
	{
		path: "/",
		element: <Layout />,
		children: [
			{
				path: "/events",
				element: <Events />,
			},
			{
				path: "events/:id",
				element: <NewEvent />,
			},
		],
	},
	{
		path: "login",
		element: <LoginPage />,
	},
	{ path: "*", element: <NotfoundPage /> },
]);

export default router;
