import { Link } from "@tanstack/react-router";
import { env } from "../env";

export const Home: React.FC = () => {
  return (
    <div>
      <h1>Home</h1>
      {env.baseUrl}
      <br />
      {window.location.host}
        <br />
      <Link to={"/tasks"} className="[&.active]:font-bold">
        list tasks
      </Link>
        <br />
      <p>Welcome to the Home page!</p>
    </div>
  );
};
