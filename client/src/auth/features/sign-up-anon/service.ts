import { IRepository } from "../../repository/repository";
import { SignUpAnonRequest } from "./dto";

export interface ISignUpAnonService {
  signUpAnon(payload: SignUpAnonRequest): Promise<void>;
}

export class SignUpAnonService implements ISignUpAnonService {
  constructor(private repository: IRepository) {}

  async signUpAnon(payload: SignUpAnonRequest): Promise<void> {
    try {
      await this.repository.firebaseSignInAnonymously();
      let currentUser = this.repository.firebaseGetCurrentUser();
      if (currentUser === null) throw new Error("failed to get current user");
      await this.repository.firebaseUpdateProfile(currentUser, {
        displayName: payload.displayName,
      });

      currentUser = this.repository.firebaseGetCurrentUser();
      if (currentUser === null) throw new Error("currentUser is null");
      const { token: idToken } = await currentUser.getIdTokenResult(false);

      await this.repository.backendExchangeToken(idToken);
      await this.repository.firebaseSignOut();
    } catch (error) {
      console.debug(error);
      throw error;
    }
  }
}
