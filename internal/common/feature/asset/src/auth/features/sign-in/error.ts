export class ValidationError extends Error {
  name: "Validation Error";
  issues: { field: string; message: string }[];

  constructor(issues: { field: string; message: string }[]) {
    super("Invalid payload");
    this.issues = issues;
  }
}
