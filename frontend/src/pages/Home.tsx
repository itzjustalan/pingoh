import { Link } from "@tanstack/react-router";
import { Button, Flex, Typography } from "antd";
import serverSvg from "../assets/server.svg";

export const Home: React.FC = () => {
  return (
    <div>
      <Typography.Title style={{ textAlign: "center", color: "#112D4E" }}>
        Pingoh
      </Typography.Title>

      <p style={{ textAlign: "center" }}>
        <Typography.Text>
          A self contained uptime monitoring tool for homelabs
        </Typography.Text>
      </p>

      <Flex wrap justify="center" align="center" gap="middle">
        <Button type="primary">
          <Link to={"/tasks"} className="[&.active]:font-bold">
            List All tasks
          </Link>
        </Button>
        <Button type="primary">
          <Link to={"/tasks/new"} className="[&.active]:font-bold">
            Create task
          </Link>
        </Button>
      </Flex>
      <img
        src={serverSvg}
        alt=""
        width={"70%"}
        style={{ margin: "0 auto", display: "block" }}
      />
    </div>
  );
};
