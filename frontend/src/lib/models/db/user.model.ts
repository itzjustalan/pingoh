import { UserRoles } from '$lib/user.access.controller';
import { z } from 'zod';

export type UserModel = z.infer<typeof userModelSchema>;
export const userModelSchema = z.object({
	id: z.string().length(36),
	email: z.string().max(255).email(),
	// passw: z.string().min(12).max(255),
	role: z.nativeEnum(UserRoles),
	// access: z.string(),
	access: z.string().array(), //todo: custom logic with zod refine
	createdAt: z.coerce.date(),
}).strict();