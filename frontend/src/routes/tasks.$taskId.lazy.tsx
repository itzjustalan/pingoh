import { createLazyFileRoute } from "@tanstack/react-router";
import { checkAuth } from "../lib/utils/authCheck";
import { TaskDetailsPage } from "../pages/tasks/TaskDetailsPage";

export const Route = createLazyFileRoute("/tasks/$taskId")({
  beforeLoad: checkAuth,
  component: () => <TaskDetailsPage />,
});
