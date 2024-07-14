import { Link, type LinkProps } from "@tanstack/react-router";
import { Menu } from "antd";

export const MenuNavItem = ({
  to,
  title,
  children,
  ...props
}: {
  title?: string;
  to: LinkProps["to"];
  children?: React.ReactNode;
}) => {
  return (
    <Menu.Item key={`${title}-${to}`} {...props}>
      <Link to={to} className="[&.active]:font-bold">
        {children ?? title ?? to}
      </Link>
    </Menu.Item>
  );
};
