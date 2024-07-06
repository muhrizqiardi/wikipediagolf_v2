import { ISignUpAnonService } from "./service";
import { htmx as htmxModule } from "../../../htmx";

export function handleSignUpAnon(
  signUpAnonService: ISignUpAnonService,
  after: () => void,
) {
  return (event: Event) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    const payload = {
      displayName: formData.get("displayName")?.toString() ?? "",
    };

    signUpAnonService.signUpAnon(payload).then(after);
  };
}
