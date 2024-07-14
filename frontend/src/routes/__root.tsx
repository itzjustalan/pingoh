import { Outlet, createRootRoute } from "@tanstack/react-router";
import { Layout, Menu } from "antd";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import React, { Suspense } from "react";
import { MenuNavItem } from "../components/MenuNavItem";
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
  component: () => {
    return (
      <>
        <Layout>
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
              <MenuNavItem to="/" title="Home" />
              <MenuNavItem to="/about" title="About" />
              <MenuNavItem to="/auth/signin" title="Signin" />
              <MenuNavItem to="/tasks/new" title="Create Task" />
            </Menu>
          </Sider>
        </Layout>
        <Layout
          style={{
            marginLeft: 200,
          }}
        >
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
        </Layout>
        <Suspense>
          <TanStackRouterDevtools />
        </Suspense>
      </>
    );
  },
});
