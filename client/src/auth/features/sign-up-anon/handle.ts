import { ISignUpAnonService } from "./service";
import { htmx as htmxModule } from "../../../htmx";

export function handleSignUpAnon(
  signUpAnonService: ISignUpAnonService,
  htmx: typeof htmxModule,
) {
  return (event: Event) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    const payload = {
      displayName: formData.get("displayName")?.toString() ?? "",
    };

    signUpAnonService.signUpAnon(payload).then(() => {
      htmx.ajax("POST", "/rooms").then(() => {
        htmx.ajax("GET", "/rooms", "body");
      });
    });
  };
}
