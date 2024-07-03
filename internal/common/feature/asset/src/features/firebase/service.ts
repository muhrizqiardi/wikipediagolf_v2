import {
  Auth,
  signInWithEmailAndPassword,
  signOut,
  User,
  UserCredential,
} from "@firebase/auth";

export interface IFirebaseService {
  getCurrentUser(): User | null;
  signInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential>;
  signOut(): Promise<void>;
}

export class FirebaseService implements IFirebaseService {
  constructor(private firebaseAuth: Auth) {
    firebaseAuth.setPersistence({ type: "NONE" });
  }

  getCurrentUser(): User | null {
    return this.firebaseAuth.currentUser;
  }

  async signInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential> {
    return await signInWithEmailAndPassword(this.firebaseAuth, email, password);
  }

  async signOut(): Promise<void> {
    await signOut(this.firebaseAuth);
  }
}
