import { FirebaseError } from "@firebase/app";
import { SignInService } from "./service";

export function handler(signInService: SignInService) {
  return (event: Event) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    const payload = {
      email: formData.get("email")?.toString() ?? "",
      password: formData.get("password")?.toString() ?? "",
    };
    signInService.signIn(payload).catch((error) => {
      let errorText = "";
      if (error instanceof FirebaseError) {
        switch (error.code) {
          case "auth/user-not-found":
            errorText = "Incorrect email or password";
        }
      }

      const signinAlertTemplate = document.querySelector(
        "#signinAlert",
      ) satisfies HTMLTemplateElement | null;
      if (signinAlertTemplate === null) {
        console.debug("error alert template not found");
        return;
      }

      const signinAlert = signinAlertTemplate.content.cloneNode(true);
      const signinAlertContainer = document.querySelector(
        "#signinAlertContainer",
      ) satisfies HTMLElement | null;
      if (signinAlertContainer === null) {
        console.debug("#signinAlertContainer element not found");
        return;
      }
      const signinAlertText = (signinAlert as Element).querySelector(
        "#signinAlertText",
      ) satisfies HTMLSpanElement | null;
      if (signinAlertText === null) {
        console.debug("#signinAlertText element not found");
        return;
      }
      signinAlertText.textContent = errorText;
      signinAlertContainer.replaceChildren(signinAlert);
    });
  };
}
