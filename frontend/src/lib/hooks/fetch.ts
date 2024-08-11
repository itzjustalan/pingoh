import { useState } from "react";
import type { FetchParams } from "../models/inputs/fetch";

export const useFetchParams = (params?: FetchParams) => {
  const [resource, setResource] = useState<FetchParams["r"]>(
    params?.r ?? "tasks",
  );
  const [id, setId] = useState<FetchParams["i"]>(params?.i);
  const [limit, setLimit] = useState<FetchParams["l"]>(params?.l);
  const [count, setCount] = useState<FetchParams["c"]>(params?.c);
  const [sort, setSort] = useState<FetchParams["s"]>(params?.s);
  const [filter, setFilter] = useState<FetchParams["f"]>(params?.f);

  return {
    resource,
    setResource,
    id,
    setId,
    limit,
    setLimit,
    count,
    setCount,
    sort,
    setSort,
    filter,
    setFilter,
  };
};
