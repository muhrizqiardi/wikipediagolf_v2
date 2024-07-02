import { SignInRequest, signInRequestSchema, SignInResponse } from "./schema";
import { IFirebaseService } from "../firebase/service";
import { ZodError } from "zod";
import { ValidationError } from "./error";

export interface ISignInService {
  signIn(payload: SignInRequest): Promise<SignInResponse>;
}

export class SignInService implements ISignInService {
  constructor(private firebaseService: IFirebaseService) {}

  async signIn(payload: SignInRequest) {
    try {
      const validPayload = signInRequestSchema.parse(payload);
      const { user } = await this.firebaseService.signInWithEmailAndPassword(
        validPayload.email,
        validPayload.password,
      );

      return user;
    } catch (error) {
      if (error instanceof ZodError) {
        throw new ValidationError(
          error.issues.map((issue) => ({
            field: issue.path.join(""),
            message: issue.message,
          })),
        );
      }
      throw error;
    }
  }
}
