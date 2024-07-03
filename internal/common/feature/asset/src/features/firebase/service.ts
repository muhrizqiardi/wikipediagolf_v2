import {
  Auth,
  browserSessionPersistence,
  signInWithEmailAndPassword,
  signOut,
  UserCredential,
} from "@firebase/auth";

export interface IFirebaseService {
  signInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential>;
  signOut(): Promise<void>;
}

export class FirebaseService implements IFirebaseService {
  constructor(private firebaseAuth: Auth) {
    firebaseAuth.setPersistence(browserSessionPersistence);
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
