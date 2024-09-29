import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { zodValidator } from "@tanstack/zod-form-adapter";
import { Button, Card, Flex, Spin, Typography } from "antd";
import { InputField } from "../../components/form/InputField";
import { signinInputSchema } from "../../lib/models/inputs/auth";
import { authNetwork } from "../../lib/networks/auth";

export const SigninPage = () => {
  const navigate = useNavigate({ from: "/auth/signin" });
  const signin = useMutation({
    mutationKey: ["signin"],
    mutationFn: authNetwork.signin,
    onSuccess: () => {
      const redirectUrl = new URL(
        new URLSearchParams(location.search).get("redirect") || location.origin,
      );
      navigate({
        to: redirectUrl.pathname,
      });
    },
  });
  const form = useForm({
    defaultValues: {
      email: "",
      passw: "",
    },
    onSubmit: async ({ value }) => {
      signin.mutateAsync(value);
    },
  });
  return (
    <>
      <p style={{ textAlign: "center" }}>
        <Typography.Title>Sign In</Typography.Title>
      </p>
      <Card style={{ maxWidth: "30vw", margin: "0 auto" }}>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            e.stopPropagation();
            form.handleSubmit();
          }}
        >
          <Flex vertical gap="small">
            <form.Field
              name="email"
              validatorAdapter={zodValidator()}
              validators={{
                onChange: signinInputSchema.shape.email,
              }}
              children={(field) => (
                <InputField field={field} label="Username:" />
              )}
            />
            <form.Field
              name="passw"
              validatorAdapter={zodValidator()}
              validators={{
                onChange: signinInputSchema.shape.passw,
              }}
              children={(field) => (
                <InputField field={field} label="Password:" type="password" />
              )}
            />

            <br />
            <form.Subscribe
              selector={(state) => [state.canSubmit, state.isSubmitting]}
              children={([canSubmit, isSubmitting]) => (
                <Button
                  type="primary"
                  htmlType="submit"
                  disabled={signin.isPending || !canSubmit}
                >
                  {signin.isPending || isSubmitting ? <Spin /> : "Submit"}
                </Button>
              )}
            />
          </Flex>
        </form>
      </Card>
    </>
  );
};
