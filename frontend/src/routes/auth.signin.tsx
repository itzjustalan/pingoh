import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { authNetwork } from "../lib/networks/auth";

export const Route = createFileRoute("/auth/signin")({
  component: () => <SigninPage />,
});

const SigninPage = () => {
  const navigate = useNavigate({ from: '/auth/signin' });
  const signin = useMutation({
    mutationKey: ["signin"],
    mutationFn: authNetwork.signin,
    onSuccess: () => navigate({ to: '/' })
  });
  const form = useForm({
    defaultValues: {
      email: "",
      passw: "",
    },
    onSubmit: async ({ value }) => {
      signin.mutateAsync(value)
    },
  });
  return (
    <>
      signin
      <form
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          form.handleSubmit();
        }}
      >
        <form.Field
          name="email"
          // validatorAdapter={zodValidator}
          // validators={{
          //   onBlur: signinInputSchema.shape.email,
          // }}
          // biome-ignore lint/correctness/noChildrenProp: <explanation>
          children={(field) => (
            <input
              name={field.name}
              value={field.state.value}
              onBlur={field.handleBlur}
              onChange={(e) => field.handleChange(e.target.value)}
            />
          )}
        />
        <form.Field
          name="passw"
          // validatorAdapter={zodValidator}
          // validators={{
          //   onBlur: signinInputSchema.shape.passw,
          // }}
          // biome-ignore lint/correctness/noChildrenProp: <explanation>
          children={(field) => (
            <input
              type="password"
              name={field.name}
              value={field.state.value}
              onBlur={field.handleBlur}
              onChange={(e) => field.handleChange(e.target.value)}
            />
          )}
        />
        <form.Subscribe
          selector={(state) => state.errors}
          // biome-ignore lint/correctness/noChildrenProp: <explanation>
          children={(errors) => errors.length > 0 && (
            <>
            {errors.toString()}
            </>
          )}
        />
        <button type="submit">Submit</button>
      </form>
    </>
  );
};
