import htmx from "./htmx";
import { initializeApp } from "@firebase/app";
import {
  connectAuthEmulator,
  getAuth,
  setPersistence,
  signInAnonymously,
  inMemoryPersistence,
} from "@firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyBCYr34q-Pn3WZ8l0Slvjdnn1dL5KFW5UU",
  authDomain: "wikipediagolf-auth-dev.firebaseapp.com",
  projectId: "wikipediagolf-auth-dev",
  storageBucket: "wikipediagolf-auth-dev.appspot.com",
  messagingSenderId: "407593468084",
  appId: "1:407593468084:web:5a2349cf5bebaaff8f17e1",
};
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

DEV: connectAuthEmulator(auth, "http://127.0.0.1:9099", {
  disableWarnings: true,
});

document.addEventListener("wg:joinRoomChooseNickname", (event) => {
  setPersistence(auth, inMemoryPersistence)
    .then(() => signInAnonymously(auth))
    .then((userCredential) => userCredential.user.getIdTokenResult())
    .then((idTokenResult) => {
      const url = `${window.location.origin}/sign-in`;
      const body = new URLSearchParams();
      body.set("idToken", idTokenResult?.token);

      return fetch(url, { method: "POST", body: body });
    })
    .then(() => htmx.ajax("get", "/rooms/join", event.currentTarget))
    .catch((error) => console.debug(error));
});

document.addEventListener("wg:createRoomChooseNickname", (event) => {
  setPersistence(auth, inMemoryPersistence)
    .then(() => signInAnonymously(auth))
    .then((userCredential) => userCredential.user.getIdTokenResult())
    .then((idTokenResult) => {
      const url = `${window.location.origin}/sign-in`;
      const body = new URLSearchParams();
      body.set("idToken", idTokenResult?.token);

      return fetch(url, { method: "POST", body: body });
    })
    .then(() =>
      htmx.ajax("post", "/rooms", {
        source: event.currentTarget,
        values: htmx.values(event.target),
      }),
    )
    .catch((error) => console.debug(error));
});
