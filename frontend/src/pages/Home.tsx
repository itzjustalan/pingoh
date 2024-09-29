import { Link } from "@tanstack/react-router";
import { env } from "../env";
import { Button, Flex } from "antd";
import serverSvg from "../assets/server.svg"



export const Home: React.FC = () => {
  return (

    <div>
      <h1 style={{ textAlign: "center", color: "#112D4E" }}>Pingoh</h1>
      <img src={serverSvg} alt="" />
      <p>A self contained uptime monitoring tool for homelabs</p>

      <Flex wrap justify="center" align="center" gap="middle">

        <Button type="primary"><Link to={"/tasks"} className="[&.active]:font-bold">
          List All tasks
        </Link></Button>
        <Button type="primary">Create Task</Button>
      </Flex>


    </div>
  );
};
