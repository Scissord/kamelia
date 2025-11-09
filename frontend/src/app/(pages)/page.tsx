import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  Button,
  Input,
  Label
} from "@/components";
import { DataTableDemo } from "./table";
import { Header } from "./header";

export default function Home() {

  return (
    <div className="p-2">
      <Header/>
      <DataTableDemo />
    </div>
  );
}
