import { useQuery } from "@tanstack/react-query";
import { getRouteApi } from "@tanstack/react-router";
import { Spin, Typography } from "antd";
import { tasksNetwork } from "../../lib/networks/tasks";

const route = getRouteApi("/tasks/$taskId");
export const TaskDetailsPage = () => {
  const { taskId } = route.useParams();
  const taskQuery = useQuery({
    queryKey: ["fetch", "tasks", taskId],
    queryFn: () =>
      tasksNetwork.fetch({
        i: Number(taskId),
      }),
  });
  if (taskQuery.isLoading) return <Spin />;
  if (!taskQuery.data?.length)
    return <Typography.Title level={2}>Task not found</Typography.Title>;

  return (
    <>
      <Typography.Title level={2}>
        {taskQuery.data[0].tasks.name}
      </Typography.Title>
      {taskId}
      {JSON.stringify(taskQuery.data[0])}
    </>
  );
};
