import { SignInService } from "./service";

export interface ISignInHandler {
  handle(formData: FormData): Promise<void>;
}

export class SignInHandler implements ISignInHandler {
  constructor(private signInService: SignInService) {}

  async handle() {}
}
