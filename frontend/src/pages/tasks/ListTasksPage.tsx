import { useQuery } from "@tanstack/react-query";
import { tasksNetwork } from "../../lib/networks/tasks";
import { useFetchParams } from "../../lib/hooks/fetch";
import { Spin, Table, type TableProps, Tag, Typography } from "antd";
import type { TaskModel } from "../../lib/models/db/task";

export const ListTasksPage = () => {
  const { limit, setLimit, count, setCount, sort, setSort, filter, setFilter } =
    useFetchParams({
      r: "tasks",
      l: 10,
      c: 0,
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
  const columns: TableProps<TaskModel>["columns"] = [
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
  return (
    <>
      <Typography.Title level={2}>
        Tasks {tasksQuery.isLoading && <Spin />}
      </Typography.Title>
      {tasksQuery.isLoading ? null : (
        <Table dataSource={tasksQuery.data} columns={columns} />
      )}
    </>
  );
};
