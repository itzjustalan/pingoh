import { createLazyFileRoute } from "@tanstack/react-router";
import { TaskDetailsPage } from "../pages/tasks/TaskDetailsPage";

export const Route = createLazyFileRoute("/tasks/$taskId")({
  component: () => <TaskDetailsPage />,
});
