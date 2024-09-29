import type { FieldApi } from "@tanstack/react-form";
import Input from "antd/es/input/Input";

export const InputField = ({
  field,
  label,
  id,
  type = "text",
}: {
  // biome-ignore lint/suspicious/noExplicitAny: <explanation>
  field: FieldApi<any, any, any, any>;
  type?: React.HTMLInputTypeAttribute;
  label: string;
  id?: string;
}) => {
  return (
    <>
      <label htmlFor={field.name}>{label}</label>
      <Input
        type={type}
        id={id ?? field.name}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
      />
      {/* <input */}
      {/*   type={type} */}
      {/*   id={field.name} */}
      {/*   name={field.name} */}
      {/*   value={field.state.value} */}
      {/*   onBlur={field.handleBlur} */}
      {/*   onChange={(e) => field.handleChange(e.target.value)} */}
      {/* /> */}
      {field.state.meta.isTouched && field.state.meta.errors.length ? (
        <em>{field.state.meta.errors.join(", ")}</em>
      ) : null}
      {field.state.meta.isValidating ? "Validating..." : null}
    </>
  );
};
