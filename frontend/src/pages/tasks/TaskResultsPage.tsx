import { useQuery } from "@tanstack/react-query";
import { Spin, Typography } from "antd";
import { useMemo } from "react";
import { Area, AreaChart, ResponsiveContainer, Tooltip, XAxis } from "recharts";
import type { ContentType } from "recharts/types/component/Tooltip";
import { tasksNetwork } from "../../lib/networks/tasks";

const customTooltip: ContentType<string, string> = ({ active, payload }) => {
  if (active && payload && payload.length) {
    return (
      <div
        style={{
          padding: "1rem",
          borderRadius: "1rem",
          border: "1px solid #ccc",
          backgroundColor: "white",
        }}
      >
        {payload.map((entry, index) => {
          return (
            <p
              key={`tooltip-item-${
                // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
                index
              }`}
            >
              <Typography.Text
                style={{
                  color: entry.payload.ok ? "green" : "red",
                }}
              >
                {entry.payload.ok ? "Up" : "Down"}
              </Typography.Text>
              <br />
              <Typography.Text>
                Status code: {entry.payload.code}
              </Typography.Text>
              <br />
              <Typography.Text>
                Duration: {Math.round(entry.payload.duration_ns / 1_000_000)} ms{" "}
              </Typography.Text>
              <br />
              <Typography.Text>
                Timestamp: {entry.payload.created_at}
              </Typography.Text>
            </p>
          );
        })}
      </div>
    );
  }
  return null;
};

export const TaskResultsPage = ({ taskId }: { taskId: string }) => {
  const taskResultsQuery = useQuery({
    queryKey: ["stats", "tasks", taskId],
    queryFn: () => tasksNetwork.stats(Number(taskId)),
    refetchInterval: 1000,
  });

  const chartData = useMemo(
    () =>
      taskResultsQuery?.data?.map((e, index: number) => ({
        ...e,
        index: index + 1,
        value: e.ok ? 1 : 0,
      })),
    [taskResultsQuery.data],
  );

  const greenColor = "#5FF55A";
  if (taskResultsQuery.isLoading) return <Spin />;
  if (taskResultsQuery.isError || !taskResultsQuery.data)
    return (
      <Typography.Text type="danger">Error Loading Results.</Typography.Text>
    );

  return (
    <>
      <div style={{ width: "100%", height: "30vh" }}>
        <ResponsiveContainer>
          <AreaChart data={chartData}>
            <XAxis dataKey="index" />
            {/* <YAxis dataKey="value" /> */}
            <Tooltip content={customTooltip} />
            <defs>
              <linearGradient id="gradientGreen" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor={greenColor} stopOpacity={0.8} />
                <stop offset="95%" stopColor={greenColor} stopOpacity={0.2} />
              </linearGradient>
            </defs>
            <Area
              type="monotone"
              dataKey="value"
              stroke={greenColor}
              fillOpacity={0.2}
              fill="url(#gradientGreen)"
            />
          </AreaChart>
        </ResponsiveContainer>
      </div>
    </>
  );
};
