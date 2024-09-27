import { useQuery } from "@tanstack/react-query";
import { getRouteApi } from "@tanstack/react-router";
import { Button, Spin, Typography } from "antd";
import { tasksNetwork } from "../../lib/networks/tasks";
import { authStore } from "../../lib/stores/auth";

const route = getRouteApi("/tasks/$taskId");
export const TaskDetailsPage = () => {
  const { taskId } = route.useParams();
  // const taskResults = new HttpResults(Number(taskId));
  // const { isConnected, sendMessage } = useWebhookConnection(
  //   (message: unknown) => {
  //     console.log("Received message:", message);
  //   },
  // );

  const startListening = () => {
    console.log("startListening");
    const url = `ws://${"127.0.0.1:3000"}/api/stats/task/${1}`;
    const token = authStore.getState().user?.access_token;
    const socket = new WebSocket(`${url}?token=${token}`);

    socket.onopen = () => {
      console.log("WebSocket connection opened");
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onmessage = (event) => {
      console.log("Message from server:", event.data);
    };

    socket.onclose = (event) => {
      console.log("WebSocket connection closed:", event);
    };
  };

  const taskQuery = useQuery({
    queryKey: ["fetch", "tasks", taskId],
    queryFn: () =>
      tasksNetwork.fetch({
        i: Number(taskId),
        ij: {
          id: "http_tasks.task_id",
        },
      }),
  });
  if (taskQuery.isLoading) return <Spin />;
  if (!taskQuery.data?.length)
    return <Typography.Title level={2}>Task not found</Typography.Title>;

  return (
    <>
      <Typography.Title level={2}>
        {taskId}: {taskQuery.data[0].tasks.name}
      </Typography.Title>
      <Button onClick={() => startListening()}>Start</Button>
      {/* <Button onClick={taskResults.stopListening}>Stop</Button> */}
      {/* <pre>{JSON.stringify(httpResults, null, 2)}</pre> */}
      <pre>{JSON.stringify(taskQuery.data[0], null, 2)}</pre>
      {/* <Button onClick={() => sendMessage("start")}>Start</Button> */}
      {/* <Button onClick={() => sendMessage("stop")}>Stop</Button> */}
    </>
  );
};
