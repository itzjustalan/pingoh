import { useQuery } from "@tanstack/react-query";
import { getRouteApi } from "@tanstack/react-router";
import { Button, Card, Divider, Spin, Tag, Typography } from "antd";
import { tasksNetwork } from "../../lib/networks/tasks";
// import { authStore } from "../../lib/stores/auth";

const route = getRouteApi("/tasks/$taskId");
export const TaskDetailsPage = () => {
  const { taskId } = route.useParams();

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
  if (taskQuery.isError || !taskQuery.data?.length)
    return <Typography.Title level={2}>Task not found</Typography.Title>;

  const taskData = taskQuery.data[0];
  return (
    <>
      <Card>
        <Typography.Title level={2}>{taskData.tasks.name}</Typography.Title>
        <Divider />
        <p>
          <Typography.Text type="secondary">
            {((): string =>
              new Date(taskData.tasks.created_at).toLocaleString())()}
          </Typography.Text>
        </p>
        <p>
          <Tag> {taskData.http_tasks.method}</Tag>
          <Typography.Link target="_blank" href={taskData.http_tasks.url}>
            {taskData.http_tasks.url}
          </Typography.Link>
        </p>
        <p>
          <Typography.Text>Status:</Typography.Text>{" "}
          <Tag color={taskData.tasks.active ? "green" : "red"}>
            {taskData.tasks.active ? "Active" : "Inactive"}
          </Tag>
        </p>

        <p>
          Accepted Status Codes:{" "}
          {JSON.parse(taskData.http_tasks.accepted_status_codes).map(
            (code: string, i: number) => (
              <Tag key={`${i}-${code}`}>{code}</Tag>
            ),
          )}
        </p>
        <p>
          Interval: <Tag>{taskData.tasks.interval} s</Tag>
        </p>

        <Typography.Paragraph
          ellipsis={{ rows: 3, expandable: true, symbol: "more" }}
        >
          {taskData.tasks.description}
        </Typography.Paragraph>
      </Card>

      {/* <Button onClick={taskResults.stopListening}>Stop</Button> */}
      {/* <pre>{JSON.stringify(httpResults, null, 2)}</pre> */}
      {/* <pre>{JSON.stringify(taskQuery.data[0], null, 2)}</pre> */}
      {/* <Button onClick={() => sendMessage("start")}>Start</Button> */}
      {/* <Button onClick={() => sendMessage("stop")}>Stop</Button> */}
    </>
  );
};
