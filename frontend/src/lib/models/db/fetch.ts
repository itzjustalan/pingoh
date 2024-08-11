import type { Task } from "./task";

export type FetchType = {
  users: unknown;
  tasks: Task;
  http_tasks: unknown;
  http_auths: unknown;
  http_results: unknown;
};
