import {
  Outlet,
  createRootRoute,
  redirect,
  useNavigate,
} from "@tanstack/react-router";
import { Layout, Menu, Typography } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import React, { Suspense } from "react";
import { MenuNavItem } from "../components/MenuNavItem";
import { env } from "../env";
import { authNetwork } from "../lib/networks/auth";
import { authStore } from "../lib/stores/auth";

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
  beforeLoad: ({ location }) => {
    if (location.pathname.startsWith("/auth/sign")) return;
    if (authStore.getState().user === undefined) {
      throw redirect({
        to: "/auth/signin",
        search: {
          redirect: location.href,
        },
      });
    }
  },
  component: () => {
    const navigate = useNavigate({ from: "/" });
    const { user } = authStore();

    return (
      <>
        <Layout>
          {user && (
            <Sider
              style={{
                overflow: "auto",
                height: "100vh",
                position: "fixed",
                top: 0,
                left: 0,
                bottom: 0,
              }}
            >
              <Menu theme="dark" mode="inline">
                <Typography.Title
                  style={{
                    color: "white",
                    textAlign: "center",
                  }}
                >
                  Pingoh
                </Typography.Title>
                <MenuNavItem to="/" title="Home" />

                <MenuNavItem to="/tasks" title="Tasks" />
                <MenuNavItem to="/tasks/new" title="Create Task" />
                <Menu.Item
                  key={"signout"}
                  onClick={() => {
                    authNetwork.signout();
                    navigate({ to: "/auth/signin" });
                  }}
                >
                  Sign Out
                </Menu.Item>
              </Menu>
            </Sider>
          )}
        </Layout>
        <Layout
          style={{
            marginLeft: user === undefined ? 0 : 200,
          }}
        >
          <div style={{ backgroundColor: "white" }}>
            {/* <Header /> */}
            <Content>
              <Outlet />
            </Content>
            {/* <Footer */}
            {/*   style={{ */}
            {/*     textAlign: "center", */}
            {/*   }} */}
            {/* > */}
            {/*   Pingoh © 2023 */}
            {/* </Footer> */}
          </div>
        </Layout>
        <Suspense>
          <TanStackRouterDevtools />
        </Suspense>
      </>
    );
  },
});
