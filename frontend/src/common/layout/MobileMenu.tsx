import {
	Sheet,
	SheetContent,
	SheetDescription,
	SheetFooter,
	SheetHeader,
	SheetTitle,
	SheetTrigger,
} from "@/components/ui/sheet";
import { Separator } from "@radix-ui/react-select";
import { Menu, User2 } from "lucide-react";

export const MobileMenu = () => {
	return (
		<div className="md:hidden">
			<Sheet>
				<SheetTrigger asChild>
					<Menu />
				</SheetTrigger>
				<SheetContent
					className="text-foreground p-0 min-h-screen overflow-auto"
					side={"left"}
				>
					<SheetHeader className="bg-primary py-5 px-5 text-white">
						<SheetTitle className="text-white text-4xl">MEvent</SheetTitle>
						<SheetDescription className="text-white/90">
							Make changes to your profile here. Click save when you're done.
						</SheetDescription>
					</SheetHeader>
					<div className="flex flex-col  ">
						<a href="/events" className="py-5 px-2 hover:bg-muted">
							Home
						</a>
						<Separator className="h-[1px] bg-primary/5" />
						<a href="/" className="py-5 px-2 hover:bg-muted">
							About
						</a>
						<Separator className="h-[1px] bg-primary/5" />
						<a href="/" className="py-5 px-2 hover:bg-muted">
							Contact
						</a>

						<Separator className="h-[1px] bg-primary/5" />
						<a href="/login" className="py-5 px-2 hover:bg-muted">
							Login
						</a>
						<Separator className="h-[1px] bg-primary/5" />
					</div>
					<SheetFooter className="absolute bottom-0 py-5">
						<div className="flex p-0 w-full gap-4">
							<User2 /> <span>username</span>
						</div>
					</SheetFooter>
				</SheetContent>
			</Sheet>
		</div>
	);
};
