import { Link, Outlet, createRootRoute } from "@tanstack/react-router";
import React, { Suspense } from "react";
import { env } from "../env";

const TanStackRouterDevtools = env.prod
  ? () => null // Render nothing in production
  : React.lazy(() =>
      // Lazy load in development
      import("@tanstack/router-devtools").then((res) => ({
        // For Embedded Mode
        // default: res.TanStackRouterDevtoolsPanel
        default: res.TanStackRouterDevtools,
      })),
    );

export const Route = createRootRoute({
  component: () => (
    <>
      <div className="p-2 flex gap-2">
        <Link to="/" className="[&.active]:font-bold">
          Home
        </Link>{" "}
        <Link to="/about" className="[&.active]:font-bold">
          About
        </Link>{" "}
        <Link to="/auth/signin" className="[&.active]:font-bold">
          Singin
        </Link>{" "}
      </div>
      <hr />
      <Outlet />
      <Suspense>
        <TanStackRouterDevtools />
      </Suspense>
    </>
  ),
});
