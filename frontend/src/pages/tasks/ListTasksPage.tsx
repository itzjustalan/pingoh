import { useQuery } from "@tanstack/react-query";
import { Link, useNavigate } from "@tanstack/react-router";
import { Spin, Table, type TableProps, Tag, Typography } from "antd";
import { Favicon } from "../../components/favicon";
import { useFetchParams } from "../../lib/hooks/fetch";
import type { FetchType } from "../../lib/models/db/fetch";
import { tasksNetwork } from "../../lib/networks/tasks";

export const ListTasksPage = () => {
  const navigate = useNavigate({ from: "/tasks" });
  // const { limit, setLimit, count, setCount, sort, setSort, filter, setFilter } =
  const { limit, count, sort, filter } = useFetchParams({
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
        ij: {
          id: "http_tasks.task_id",
        },
      }),
  });

  const columns: TableProps<FetchType>["columns"] = [
    {
      title: "ID",
      dataIndex: ["tasks", "id"],
      key: "id",
    },
    {
      title: "Name",
      dataIndex: ["tasks", "name"],
      key: "name",
      render: (name, record) => (
        <Typography.Text
          style={{
            fontSize: 28,
          }}
        >
          <Favicon url={record.http_tasks.url} /> {name}
        </Typography.Text>
      ),
    },
    {
      title: "Status",
      dataIndex: ["tasks", "active"],
      key: "active",
      render: (active: boolean) => (
        <Tag color={active ? "green" : "red"}>
          {active ? "Active" : "Inactive"}
        </Tag>
      ),
    },
    {
      title: "URL",
      dataIndex: ["http_tasks", "url"],
      key: "url",
      render: (url) => (
        <Typography.Link href={url} target="_blank">
          {url}
        </Typography.Link>
      ),
    },
  ];

  const nodata = () => (
    <Typography.Text type="secondary">
      You have not created any tasks yet. <br />
      <Link to="/tasks/new">Create a new task</Link>
    </Typography.Text>
  );

  const table = () => (
    <Table
      dataSource={tasksQuery.data}
      columns={columns}
      onRow={(row) => {
        return {
          onClick: () => {
            navigate({
              to: "/tasks/$taskId",
              params: { taskId: row.tasks.id.toString() },
            });
          },
        };
      }}
    />
  );

  return (
    <>
      <Typography.Title level={2}>
        Tasks {tasksQuery.isLoading && <Spin />}
      </Typography.Title>
      {(() => {
        if (tasksQuery.isLoading) return null;
        if (tasksQuery.data?.length === 0) {
          return nodata();
        }
        return table();
      })()}
    </>
  );
};
