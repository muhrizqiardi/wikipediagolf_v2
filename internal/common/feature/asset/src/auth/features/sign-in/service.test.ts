import { describe, expect, it, vi } from "vitest";
import { IFirebaseService } from "../firebase/service";
import { SignInRequest } from "./schema";
import { SignInService } from "./service";
import { ValidationError } from "./error";
import { User, UserCredential } from "@firebase/auth";

describe("SignInService", () => {
  it("should handle validation failure", () => {
    let firebaseService: IFirebaseService = {} as unknown as IFirebaseService;
    const signInService = new SignInService(firebaseService);
    const payload: SignInRequest = {
      email: "invalid email",
      password: "validPassword123?",
    };

    expect(() => signInService.SignIn(payload)).toThrowError(ValidationError);
  });
  it("should sign in via FirebaseService", async () => {
    let mockFirebaseService: IFirebaseService = {
      signInWithEmailAndPassword: async (_email, _password) =>
        ({}) as unknown as UserCredential,
    };

    vi.spyOn(
      mockFirebaseService,
      "signInWithEmailAndPassword",
    ).mockImplementationOnce(async (email, _) => {
      return {
        providerId: "mockProviderId",
        operationType: "signIn",
        user: {
          email: email,
          displayName: "display name",
          emailVerified: true,
        } as unknown as User,
      };
    });

    const signInService = new SignInService(mockFirebaseService);
    expect(
      await signInService.SignIn({
        email: "mock@example.com",
        password: "strong_Password321",
      }),
    ).toEqual({
      email: "mock@example.com",
      displayName: "display name",
      emailVerified: true,
    } as unknown as User);
  });
});
