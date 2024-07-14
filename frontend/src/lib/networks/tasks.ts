import { type TaskModel, taskModelSchema } from "../models/db/task";
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
}

export const tasksNetwork = new TasksNetwork();
