import { createLazyFileRoute } from "@tanstack/react-router";
import { CreateTaskPage } from "../pages/CreateTaskPage";

export const Route = createLazyFileRoute("/tasks/new")({
  component: () => <CreateTaskPage />,
});
