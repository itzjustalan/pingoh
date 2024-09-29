import { createLazyFileRoute } from "@tanstack/react-router";
import { checkAuth } from "../lib/utils/authCheck";
import { Home } from "../pages/Home";

export const Route = createLazyFileRoute("/")({
  beforeLoad: checkAuth,
  component: () => <Home />,
});
