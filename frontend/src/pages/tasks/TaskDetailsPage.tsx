import { useMutation, useQuery } from "@tanstack/react-query";
import { getRouteApi, useNavigate } from "@tanstack/react-router";
import {
  Button,
  Card,
  Divider,
  Flex,
  Popconfirm,
  Spin,
  Tag,
  Typography,
} from "antd";
import { Favicon } from "../../components/favicon";
import { tasksNetwork } from "../../lib/networks/tasks";

const route = getRouteApi("/tasks/$taskId");
export const TaskDetailsPage = () => {
  const { taskId } = route.useParams();
  const navigate = useNavigate({ from: "/tasks/$taskId" });

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

  const toggleTask = useMutation({
    mutationKey: ["task", taskId, "toggle"],
    mutationFn: tasksNetwork.toggle,
    onSuccess: () => {
      taskQuery.refetch();
    },
  });

  const deleteTask = useMutation({
    mutationKey: ["task", taskId, "delete"],
    mutationFn: tasksNetwork.delete,
    onSuccess: () => {
      navigate({
        to: "/tasks",
      });
    },
  });

  if (taskQuery.isLoading) return <Spin />;
  if (taskQuery.isError || !taskQuery.data?.length)
    return <Typography.Title level={2}>Task not found</Typography.Title>;

  const greenColor = "#5FF55A";
  const taskData = taskQuery.data[0];
  return (
    <>
      <Card>
        <Flex justify="space-between" align="center">
          <Typography.Title level={1}>
            <Favicon url={taskData.http_tasks.url} /> {taskData.tasks.name}
          </Typography.Title>
          <p>
            <Button
              htmlType="button"
              onClick={() => toggleTask.mutateAsync(taskData.tasks.id)}
              style={{
                color: !taskData.tasks.active ? greenColor : "red",
                borderColor: !taskData.tasks.active ? greenColor : "red",
              }}
            >
              {!taskData.tasks.active ? "Activate" : "Deactivate"}
            </Button>{" "}
            <Popconfirm
              title="Delete the task"
              description="Are you sure you want to delete this task?"
              onConfirm={async () =>
                await deleteTask.mutateAsync(taskData.tasks.id)
              }
            >
              <Button htmlType="button" type="primary" danger>
                Delete
              </Button>
            </Popconfirm>
          </p>
        </Flex>
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
