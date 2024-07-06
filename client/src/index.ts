import { initializeApp } from "@firebase/app";
import { connectAuthEmulator, getAuth } from "@firebase/auth";
import { htmx as htmxModule } from "./htmx";
import { Repository } from "./auth/repository/repository";
import { SignInService } from "./auth/features/sign-in/service";
import { handleSignIn } from "./auth/features/sign-in/handle";
import { handleSignUpAnon } from "./auth/features/sign-up-anon/handle";
import { SignUpAnonService } from "./auth/features/sign-up-anon/service";

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

const repository = new Repository(auth);
const signInService = new SignInService(repository);
const signUpAnonService = new SignUpAnonService(repository);
const signInHandler = handleSignIn(signInService);
const signUpAnonCreateRoomHandler = handleSignUpAnon(signUpAnonService, () => {
  window.htmx.ajax("POST", "/rooms").then(() => {
    window.htmx.ajax("GET", "/rooms", "body");
  });
});
const signUpAnonJoinRoomHandler = handleSignUpAnon(signUpAnonService, () => {
  window.htmx
    .ajax("POST", "/rooms/join", {
      source: "#joinNicknameModalForm",
      values: window.htmx.values(window.htmx.find("#joinNicknameModalForm")!),
    })
    .then(() => {
      window.htmx.ajax("GET", "/rooms", "body");
    });
});

function addEventHandlers() {
  const signinFormEl = document.querySelector("#signin");
  signinFormEl?.addEventListener("submit", signInHandler);

  const chooseNicknameModalFormEl =
    document.querySelector("#nicknameModalForm");
  chooseNicknameModalFormEl?.addEventListener(
    "submit",
    signUpAnonCreateRoomHandler,
  );

  const chooseNicknameModalJoinFormEl = document.querySelector(
    "#joinNicknameModalForm",
  );
  chooseNicknameModalJoinFormEl?.addEventListener(
    "submit",
    signUpAnonJoinRoomHandler,
  );

  const signupFormEl = document.querySelector("#signup");
  if (signupFormEl !== null) {
    window.htmx.on("#signup", "htmx:afterRequest", (evt) => {
      const htmxEvt = evt as Event & {
        detail?: {
          elt?: Element;
          xhr?: XMLHttpRequest;
          target?: EventTarget;
          requestConfig?: object;
          successful?: boolean;
          failed?: boolean;
        };
      };
      if (htmxEvt.detail === undefined || htmxEvt.detail === null) return;
      if (htmxEvt.detail.xhr === undefined || htmxEvt.detail.xhr === null)
        return;
      if (htmxEvt.detail.xhr.status === 201) {
        signInHandler(evt);
      }
    });
  }
}
window.htmx.onLoad(() => addEventHandlers());
