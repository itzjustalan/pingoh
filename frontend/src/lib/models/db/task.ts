import { z } from "zod";

const httpAllowedEncodings = ["none", "text", "html", "form", "json", "xml"];

const httpAllowedMethods = [
  "GET",
  "POST",
  "PUT",
  "PATCH",
  "DELETE",
  "HEAD",
  "OPTIONS",
];

const zodStringEquals = (value: string) =>
  z.string().refine((v) => v === value);
const zodStringOneOf = (values: string[]) =>
  z.string().refine((v) => values.includes(v));
const zodUniqueArray = (schema: z.ZodTypeAny) =>
  z.array(schema).refine((v) => new Set(v).size === v.length);

export const createTaskSchema = z
  .object({
    name: z.string(),
    type: zodStringEquals("http"),
    repeat: z.boolean(),
    active: z.boolean(),
    interval: z.number().gte(1), // seconds
    description: z.string(),
    tags: zodUniqueArray(z.string()),
    http: z.object({
      method: zodStringOneOf(httpAllowedMethods),
      url: z.string().url(),
      encoding: zodStringOneOf(httpAllowedEncodings),
      headers: z.object({}),
      retries: z.number().gte(0),
      timeout: z.number().gte(0), // seconds
      accepted_status_codes: z.number().gte(100).lt(600).array(),
      auth_method: zodStringOneOf(["none", "basic", "oauth2"]),
    }),
  })
  .strict();
export type NewTask = z.infer<typeof createTaskSchema>;
export type Task = NewTask & { id: number; created_at: string };
