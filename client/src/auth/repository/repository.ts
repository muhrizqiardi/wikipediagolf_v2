import {
  Auth,
  signInAnonymously,
  signInWithEmailAndPassword,
  signOut,
  updateProfile,
  User,
  UserCredential,
} from "@firebase/auth";

export interface IRepository {
  backendExchangeToken(idToken: string): Promise<void>;
  firebaseGetCurrentUser(): User | null;
  firebaseUpdateProfile(
    oldUser: User,
    newUser: { displayName?: string; photoURL?: string },
  ): Promise<void>;
  firebaseSignInAnonymously(): Promise<UserCredential>;
  firebaseSignInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential>;
  firebaseSignOut(): Promise<void>;
}

export class Repository implements IRepository {
  constructor(private firebaseAuth: Auth) {
    firebaseAuth.setPersistence({ type: "NONE" });
  }

  async backendExchangeToken(idToken: string): Promise<void> {
    try {
      const tokenExchangeReq = new URLSearchParams();
      tokenExchangeReq.set("idToken", idToken);
      const tokenExchangeURL = new URL("/sign-in", window.location.origin);
      const tokenExchangeResult = await fetch(tokenExchangeURL, {
        method: "POST",
        body: tokenExchangeReq,
      });
      if (!tokenExchangeResult.ok) throw new Error("failed to sign in");
    } catch (error) {
      throw error;
    }
  }

  async firebaseUpdateProfile(
    oldUser: User,
    newUser: { displayName?: string; photoURL?: string },
  ): Promise<void> {
    return await updateProfile(oldUser, newUser);
  }

  firebaseGetCurrentUser(): User | null {
    return this.firebaseAuth.currentUser;
  }

  async firebaseSignInAnonymously(): Promise<UserCredential> {
    return await signInAnonymously(this.firebaseAuth);
  }

  async firebaseSignInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential> {
    return await signInWithEmailAndPassword(this.firebaseAuth, email, password);
  }

  async firebaseSignOut(): Promise<void> {
    await signOut(this.firebaseAuth);
  }
}
