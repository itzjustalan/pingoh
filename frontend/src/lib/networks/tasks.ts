import type { FetchType } from "../models/db/fetch";
import { type NewTask, createTaskSchema } from "../models/db/task";
import {
  type FetchParams,
  fetchQuerySchema,
  qStringFromParams,
} from "../models/inputs/fetch";
import backendApi from "./apis/backend";

export type HttpTaskResult = {
  task_id: number;
  code: number;
  ok: boolean;
  duration_ns: number;
  created_at: string;
};

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
  stats = async (taskId: number) => {
    const res = await backendApi.get<HttpTaskResult[]>(`/stats/task/${taskId}`);
    return res.data ?? [];
  };
  activate = async (taskId: number): Promise<void> => {
    const res = await backendApi.post(`/tasks/${taskId}/activate`);
    return res.data;
  };
  deactivate = async (taskId: number): Promise<void> => {
    const res = await backendApi.post(`/tasks/${taskId}/deactivate`);
    return res.data;
  };
  toggle = async (taskId: number): Promise<void> => {
    const res = await backendApi.get(`/tasks/${taskId}/toggle`);
    return res.data;
  };
  drop = async (taskId: number): Promise<void> => {
    const res = await backendApi.delete(`/tasks/${taskId}`);
    return res.data;
  };
  subscribe = (_taskId: number) => {
    // const ws = new WebSocket("ws://localhost:8080/ws");
    // const ws = new WebSocket(
    //   `ws://${env.backend.baseUrl}/api/stats/ws/task/${taskId}?token=${authStore.getState().user?.access_token}`,
    // );
  };
}

export const tasksNetwork = new TasksNetwork();
