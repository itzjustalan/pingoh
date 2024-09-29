import { createLazyFileRoute } from "@tanstack/react-router";
import { TaskDetailsPage } from "../pages/tasks/TaskDetailsPage";
import { checkAuth } from "../lib/utils/authCheck";

export const Route = createLazyFileRoute("/tasks/$taskId")({
  beforeLoad: checkAuth,
  component: () => <TaskDetailsPage />,
});
