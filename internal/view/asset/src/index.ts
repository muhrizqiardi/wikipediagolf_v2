import { initializeApp } from "@firebase/app";
import { connectAuthEmulator, getAuth } from "@firebase/auth";
import { SignInService } from "./features/sign-in/service";
import { FirebaseService } from "./features/firebase/service";
import { handler } from "./features/sign-in/handle";
import { htmx as htmxModule } from "./htmx";

declare const window: Window &
  typeof globalThis & {
    htmx: typeof htmxModule;
  };

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

window.htmx = htmxModule;
if (window.htmx === null || window.htmx === undefined)
  console.error("window.htmx is undefined or null");

DEV: connectAuthEmulator(auth, "http://127.0.0.1:9099");

const firebaseService = new FirebaseService(auth);
const signInService = new SignInService(firebaseService);
const signInHandler = handler(signInService);

function addEventHandlers() {
  const signinFormEl = document.querySelector("#signin");
  if (signinFormEl === null) return null;

  signinFormEl.addEventListener("submit", signInHandler);
  window.htmx.on("htmx:beforeRequest", () => {
    console.debug("bebek");
  });
}
addEventHandlers();
