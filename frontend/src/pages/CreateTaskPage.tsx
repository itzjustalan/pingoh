import { type FieldApi, useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { tasksNetwork } from "../lib/networks/tasks";

export const CreateTaskPage = () => {
  const createTask = useMutation({
    mutationKey: ["create", "task"],
    mutationFn: tasksNetwork.create,
  });
  const form = useForm({
    defaultValues: {
      name: "",
      type: "http",
      repeat: true,
      active: true,
      interval: 60,
      description: "",
      tags: [],
      http: {
        method: "GET",
        url: "",
        encoding: "none",
        headers: {},
        retries: 3,
        timeout: 10, // seconds
        accepted_status_codes: [200],
        auth_method: "none",
      },
    },
    onSubmit: async ({ value }) => {
      console.log(value, value.http);
      createTask.mutateAsync(value);
    },
  });
  return (
    <>
      new task
      <form
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          form.handleSubmit();
        }}
      >
        <form.Field
          name="name"
          children={(field) => <InputField field={field} label="Name:" />}
        />
        <br />
        <form.Field
          name="interval"
          children={(field) => <InputField field={field} type="number" label="Interval:" />}
        />
        <br />
        <form.Field
          name="description"
          children={(field) => <InputField field={field} label="Description:" />}
        />
        <br />
        <form.Field
          mode="array"
          name="http.url"
          children={(field) => <InputField field={field} type="url" label="Url:" />}
        />
        <br />
        <button type="submit">Submit</button>
      </form>
    </>
  );
};

function InputField({
  field,
  label,
  type = "text",
}: {
  // biome-ignore lint/suspicious/noExplicitAny: <explanation>
  field: FieldApi<any, any, any, any>;
  type?: React.HTMLInputTypeAttribute;
  label: string;
}) {
  return (
    <>
      <label htmlFor={field.name}>{label}</label>
      <input
        type={type}
        id={field.name}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
      />
      {field.state.meta.isTouched && field.state.meta.errors.length ? (
        <em>{field.state.meta.errors.join(", ")}</em>
      ) : null}
      {field.state.meta.isValidating ? "Validating..." : null}
    </>
  );
}
