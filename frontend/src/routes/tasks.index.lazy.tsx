import { createLazyFileRoute } from "@tanstack/react-router";
import { checkAuth } from "../lib/utils/authCheck";
import { ListTasksPage } from "../pages/tasks/ListTasksPage";

export const Route = createLazyFileRoute("/tasks/")({
  beforeLoad: checkAuth,
  component: () => <ListTasksPage />,
});
