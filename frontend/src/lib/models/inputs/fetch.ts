import { z } from "zod";

export const fetchQuerySchema = z.object({
  r: z.enum(["tasks", "users"]),
  i: z.coerce.number().gte(1).optional(),
  l: z.coerce.number().gte(1).optional(),
  c: z.coerce.number().gte(1).optional(),
  s: z.record(z.string(), z.enum(["a", "d"])).optional(),
  f: z.record(z.string(), z.coerce.string()).optional(),
  ij: z.record(z.string(), z.coerce.string()).optional(),
});

export type FetchParams = z.infer<typeof fetchQuerySchema>;

export const qStringFromParams = (params: FetchParams): string => {
  const q = new URLSearchParams();
  for (const key in params) {
    const val = params[key as keyof typeof params] ?? "";
    if (typeof val === "object") {
      for (const subkey in val) {
        q.append(
          `${key}[${subkey}]`,
          val[subkey as keyof typeof val]?.toString() ?? "",
        );
      }
    } else {
      q.append(key, val.toString());
    }
  }
  return q.toString();
};
