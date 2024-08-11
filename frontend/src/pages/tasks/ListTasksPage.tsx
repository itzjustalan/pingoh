import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { Button, Spin, Table, type TableProps, Tag, Typography } from "antd";
import { useFetchParams } from "../../lib/hooks/fetch";
import type { Task } from "../../lib/models/db/task";
import { tasksNetwork } from "../../lib/networks/tasks";

export const ListTasksPage = () => {
  const navigate = useNavigate({ from: "/tasks" });
  const { limit, setLimit, count, setCount, sort, setSort, filter, setFilter } =
    useFetchParams({
      r: "tasks",
      l: 10,
      c: 1,
    });
  const tasksQuery = useQuery({
    queryKey: ["fetch", "tasks", limit, count, sort, filter],
    queryFn: () =>
      tasksNetwork.fetch({
        l: limit,
        c: count,
        s: sort,
        f: filter,
      }),
  });

  const columns: TableProps<Task>["columns"] = [
    {
      title: "ID",
      dataIndex: "id",
      key: "id",
    },
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Status",
      dataIndex: "active",
      key: "active",
      render: (active: boolean) => (
        <Tag color={active ? "green" : "red"}>
          {active ? "Active" : "Inactive"}
        </Tag>
      ),
    },
  ];
  console.log(tasksQuery.error);

  return (
    <>
      <Typography.Title level={2}>
        Tasks {tasksQuery.isLoading && <Spin />}
      </Typography.Title>
      <Button
        htmlType="button"
        onClick={() => {
          setSort({ id: "d" });
        }}
      >
        sort
      </Button>
      {tasksQuery.isLoading ? null : (
        <Table
          dataSource={tasksQuery.data}
          columns={columns}
          onRow={(row) => {
            return {
              onClick: () => {
                navigate({
                  to: "/tasks/$taskId",
                  params: { taskId: row.id.toString() },
                });
              },
            };
          }}
        />
      )}
    </>
  );
};
