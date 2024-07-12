import { DataTable } from "@/components/custom/DataTable";
import { PageHaeder } from "@/components/custom/PageHaeder";
import { TableContainer } from "@/components/custom/TableContainer";
import { useQuery } from "@tanstack/react-query";
import { Edit3, Loader, Trash } from "lucide-react";
import { useMemo } from "react";
import { getEvents } from "./fetchers";
import { Button } from "@/components/ui/button";
// biome-ignore lint/style/useImportType: <explanation>
import { ColumnDef } from "@tanstack/react-table";
// biome-ignore lint/style/useImportType: <explanation>
import { TEvent } from "@/lib/types/events";
import apiClient from "@/common/api";
import { Link } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
export const Events = (_props: { pageName?: string }) => {
	const { isLoading, data } = useQuery({
		queryKey: ["events"],
		queryFn: () => {
			return getEvents();
		},
	});
	const columns = useMemo<ColumnDef<TEvent>[]>(() => {
		return [
			{
				accessorKey: "id",
				header: "N",
			},
			{
				accessorKey: "images",
				header: "image",
				cell(props) {
					return (
						<Avatar>
							<AvatarImage
								src={
									"http://127.0.0.1:8080" +
									props.row.original?.images[0].replace("./storage", "")
								}
							/>
							<AvatarFallback>
								{props.row.original.name.slice(0, 2)}
							</AvatarFallback>
						</Avatar>
					);
				},
			},
			{
				accessorKey: "name",
				header: "Name",
			},
			{
				accessorKey: "description",
				header: "Description",
			},
			{
				accessorKey: "date",
				header: "Date",
			},
			{
				accessorKey: "tickets",
				header: "Tickets",
			},
			{
				accessorKey: "avialable_tickets",
				header: "Avialable Tickets",
				cell(props) {
					return (
						<span className="bg-primary rounded-md px-2 py-1 text-white font-bold">
							{props.row.original.avialable_tickets}
						</span>
					);
				},
			},
			{
				header: "Actions",
				cell(props) {
					return (
						<div className="flex gap-2">
							<Button size={"icon"} variant={"secondary"}>
								<Link to={`/events/${props.row.original.id}`}>
									<Edit3 />
								</Link>
							</Button>
							<Button
								size={"icon"}
								variant={"destructive"}
								onClick={async () => {
									await apiClient.delete(`/events/${props.row.original.id}`);
								}}
							>
								<Trash />
							</Button>
						</div>
					);
				},
			},
		];
	}, []);
	if (isLoading) {
		return <Loader className="animate-spin" size={64} />;
	}
	return (
		<div className="animate-in transition-all">
			<PageHaeder title="Events" newPage="new" />

			<TableContainer>
				<DataTable data={data || []} columns={columns} />
			</TableContainer>
		</div>
	);
};
