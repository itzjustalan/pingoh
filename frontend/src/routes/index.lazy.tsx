import { createLazyFileRoute } from "@tanstack/react-router";
import { Home } from "../pages/Home";
import { checkAuth } from "../lib/utils/authCheck";

export const Route = createLazyFileRoute("/")({
  beforeLoad: checkAuth,
  component: () => <Home />,
});
