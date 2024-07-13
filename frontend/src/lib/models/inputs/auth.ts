import { z } from 'zod';
// import { userModelSchema } from '../db/user.model';

export type SignupInput = z.infer<typeof signupInputSchema>;
export type SigninInput = z.infer<typeof signinInputSchema>;
export const signupInputSchema = z
	.object({
		name: z.string(),
		email: z.string().max(255).email(),
		passw: z.string().min(8).max(50)
	})
	.strict();
export const signinInputSchema = z
	.object({
		email: z.string().max(255).email(),
		passw: z.string().min(8).max(50)
	})
	.strict();

// export type UserInput = z.infer<typeof userInputSchema>;
// export const userInputSchema = userModelSchema
// 	.omit({
// 		id: true,
// 		createdAt: true,
// 		updatedAt: true
// 	})
// 	.strict();
