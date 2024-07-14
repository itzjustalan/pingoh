import { env } from "../env";

export const Home: React.FC = () => {
  return (
    <div>
      <h1>Home</h1>
      {env.baseUrl}
      <br />
      {window.location.host}
      <p>Welcome to the Home page!</p>
    </div>
  );
};
