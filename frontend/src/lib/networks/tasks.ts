import { type TaskModel, taskModelSchema } from "../models/db/task";
import {
  fetchQuerySchema,
  type FetchParams,
  qStringFromParams,
} from "../models/inputs/fetch";
import { sleep } from "../utils";
import backendApi from "./apis/backend";

class TasksNetwork {
  // create = async <T extends TaskModel>(data: T): Promise<T> => {
  //   taskModelSchema.parse(data);
  //   const res = await backendApi.post<T>("/tasks", data);
  //   return res.data;
  // };
  create = async (data: TaskModel): Promise<TaskModel> => {
    taskModelSchema.parse(data);
    const res = await backendApi.post<TaskModel>("/tasks", data);
    return res.data;
  };
  fetch = async (params?: Omit<FetchParams, "r">): Promise<TaskModel[]> => {
    await sleep(1000 * 2);
    const q = qStringFromParams({ ...params, r: "tasks" });
    const res = await backendApi.get<TaskModel[]>(`/shared/fetch?${q}`);
    return res.data ?? [];
  };
}

export const tasksNetwork = new TasksNetwork();
