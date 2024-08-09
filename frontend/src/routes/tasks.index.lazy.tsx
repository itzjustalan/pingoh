import { createLazyFileRoute } from "@tanstack/react-router";
import { ListTasksPage } from "../pages/tasks/ListTasksPage";

export const Route = createLazyFileRoute("/tasks/")({
  component: () => <ListTasksPage />,
});
