import { useEffect, useState } from "react";
import { env } from "../../env";
import { authStore } from "../stores/auth";

export const useHttpResults = (taskId: number) => {
  const [httpResults, setHttpResults] = useState();
  const url = `ws://${env.backend.baseUrl}/api/stats/task/${taskId}`;
  const token = authStore.getState().user?.access_token;
  const ws = new WebSocket(`${url}?token=${token}`);

  useEffect(() => {
    ws.onopen = () => {
      console.log("connected");
    };

    ws.onmessage = (event) => {
      console.log(event);
      // setResults(JSON.parse(event.data));
    };

    return () => {
      ws.close();
    };
  }, [ws]);

  return {
    httpResults,
    startListening: () => {
      ws.send("start");
    },
    stopListening: () => {
      ws.send("stop");
    },
  };
};
