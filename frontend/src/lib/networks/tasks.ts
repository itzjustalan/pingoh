import { type NewTask, type Task, createTaskSchema } from "../models/db/task";
import {
  type FetchParams,
  fetchQuerySchema,
  qStringFromParams,
} from "../models/inputs/fetch";
import { sleep } from "../utils";
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
  fetch = async (params?: Omit<FetchParams, "r">): Promise<Task[]> => {
    const qParams = fetchQuerySchema.parse({ ...params, r: "tasks" });
    const q = qStringFromParams(qParams);
    const res = await backendApi.get<Task[]>(`/shared/fetch?${q}`);
    return res.data ?? [];
  };
}

export const tasksNetwork = new TasksNetwork();
