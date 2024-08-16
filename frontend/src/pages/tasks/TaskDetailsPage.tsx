import { useQuery } from "@tanstack/react-query";
import { getRouteApi } from "@tanstack/react-router";
import { Spin, Typography } from "antd";
import { tasksNetwork } from "../../lib/networks/tasks";
import { useHttpResults } from "../../lib/hooks/stats";

const route = getRouteApi("/tasks/$taskId");
export const TaskDetailsPage = () => {
  const { taskId } = route.useParams();
  const { httpResults, startListening, stopListening } = useHttpResults(
    Number(taskId),
  );
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
      <pre>{JSON.stringify(httpResults, null, 2)}</pre>
      <pre>{JSON.stringify(taskQuery.data[0], null, 2)}</pre>
    </>
  );
};
