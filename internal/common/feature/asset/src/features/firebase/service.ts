import {
  Auth,
  signInWithEmailAndPassword,
  UserCredential,
} from "@firebase/auth";

export interface IFirebaseService {
  signInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential>;
}

export class FirebaseService implements IFirebaseService {
  constructor(private firebaseAuth: Auth) {}

  async signInWithEmailAndPassword(
    email: string,
    password: string,
  ): Promise<UserCredential> {
    return await signInWithEmailAndPassword(this.firebaseAuth, email, password);
  }
}
