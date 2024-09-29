import { createLazyFileRoute } from "@tanstack/react-router";
import { checkAuth } from "../lib/utils/authCheck";
import { CreateTaskPage } from "../pages/tasks/CreateTaskPage";

export const Route = createLazyFileRoute("/tasks/new")({
  beforeLoad: checkAuth,
  component: () => <CreateTaskPage />,
});
