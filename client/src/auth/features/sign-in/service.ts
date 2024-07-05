import { SignInRequest, signInRequestSchema, SignInResponse } from "./schema";
import { ZodError } from "zod";
import { ValidationError } from "./error";
import { IRepository } from "../../repository/repository";

export interface ISignInService {
  signIn(payload: SignInRequest): Promise<SignInResponse>;
}

export class SignInService implements ISignInService {
  constructor(private repository: IRepository) {}

  async signIn(payload: SignInRequest) {
    try {
      const validPayload = signInRequestSchema.parse(payload);
      const { user } = await this.repository.firebaseSignInWithEmailAndPassword(
        validPayload.email,
        validPayload.password,
      );
      const currentUser = this.repository.firebaseGetCurrentUser();
      if (currentUser === null) throw new Error("currentUser is null");
      const { token: idToken } = await currentUser.getIdTokenResult(false);

      await this.repository.backendExchangeToken(idToken);
      await this.repository.firebaseSignOut();

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
