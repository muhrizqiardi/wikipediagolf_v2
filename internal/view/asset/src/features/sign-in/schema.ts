import { User } from "@firebase/auth";
import z from "zod";

export const signInRequestSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export interface SignInRequest extends z.infer<typeof signInRequestSchema> {}

export interface SignInResponse extends User {}
