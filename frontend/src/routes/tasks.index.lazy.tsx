import { createLazyFileRoute } from "@tanstack/react-router";
import { ListTasksPage } from "../pages/tasks/ListTasksPage";
import { checkAuth } from "../lib/utils/authCheck";

export const Route = createLazyFileRoute("/tasks/")({
  beforeLoad: checkAuth,
  component: () => <ListTasksPage />,
});
