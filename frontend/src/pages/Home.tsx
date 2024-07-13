import { env } from "../env";
import { authNetwork } from "../lib/networks/auth";

export const Home: React.FC = () => {
  const test = async () => {
    const res = await authNetwork.signin({
      email: 'admin01@mail.com',
      passw: 'qwe123!@#',
    })
    console.log(res)
  }
  return (
    <div>
      <h1 onClick={test}>Home</h1>
      {env.baseUrl}
      <br />
      {window.location.host}
      <p>Welcome to the Home page!</p>
    </div>
  );
};
