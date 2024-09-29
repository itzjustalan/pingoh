import { createLazyFileRoute } from "@tanstack/react-router";
import { checkAuth } from "../lib/utils/authCheck";

export const Route = createLazyFileRoute("/about")({
  beforeLoad: checkAuth,
  component: () => <div>Hello /about!</div>,
});
