import { HttpTask } from "./http_task";
import type { Task } from "./task";

export type FetchType = {
  users: unknown;
  tasks: Task;
  http_tasks: HttpTask;
  http_auths: unknown;
  http_results: unknown;
};
