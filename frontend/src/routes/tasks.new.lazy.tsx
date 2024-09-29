import { createLazyFileRoute } from "@tanstack/react-router";
import { CreateTaskPage } from "../pages/tasks/CreateTaskPage";
import { checkAuth } from "../lib/utils/authCheck";

export const Route = createLazyFileRoute("/tasks/new")({
  beforeLoad: checkAuth,
  component: () => <CreateTaskPage />,
});
