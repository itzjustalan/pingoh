import { z } from "zod";

export const UserRoles = {
  Admin: "admin",
  Guest: "guest",
  User: "user",
} as const;

export type UserModel = z.infer<typeof userModelSchema>;
export const userModelSchema = z
  .object({
    id: z.string().length(36),
    name: z.string().max(255),
    email: z.string().max(255).email(),
    role: z.nativeEnum(UserRoles),
    access: z.string().array(), //todo: custom logic with zod refine
    createdAt: z.coerce.date(),
  })
  .strict();
