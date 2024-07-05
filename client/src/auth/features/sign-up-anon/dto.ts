import { z } from "zod";

export const signUpAnonRequestSchema = z.object({
  displayName: z.string(),
});

export interface SignUpAnonRequest
  extends z.infer<typeof signUpAnonRequestSchema> {}
