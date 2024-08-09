import { createLazyFileRoute } from "@tanstack/react-router";
import { CreateTaskPage } from "../pages/tasks/CreateTaskPage";

export const Route = createLazyFileRoute("/tasks/new")({
  component: () => <CreateTaskPage />,
});
