import type { FetchType } from "../models/db/fetch";
import { type NewTask, createTaskSchema } from "../models/db/task";
import {
  type FetchParams,
  fetchQuerySchema,
  qStringFromParams,
} from "../models/inputs/fetch";
import backendApi from "./apis/backend";

class TasksNetwork {
  // dcreate = async <T extends NewTask>(data: T): Promise<T> => {
  //   createTaskModelSchema.parse(data);
  //   const res = await backendApi.post<T>("/tasks", data);
  //   return res.data;
  // };
  create = async (data: NewTask): Promise<NewTask> => {
    const task = createTaskSchema.parse(data);
    const res = await backendApi.post<NewTask>("/tasks", task);
    return res.data;
  };
  fetch = async (params?: Omit<FetchParams, "r">) => {
    const qParams = fetchQuerySchema.parse({ ...params, r: "tasks" });
    const q = qStringFromParams(qParams);
    const res = await backendApi.get<FetchType[]>(`/shared/fetch?${q}`);
    return res.data ?? [];
  };
  subscribe = (_taskId: number) => {
    // const ws = new WebSocket("ws://localhost:8080/ws");
    // const ws = new WebSocket(
    //   `ws://${env.backend.baseUrl}/api/stats/ws/task/${taskId}?token=${authStore.getState().user?.access_token}`,
    // );
  };
}

export const tasksNetwork = new TasksNetwork();
