import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { authNetwork } from "../../lib/networks/auth";
import { Button, Card, Flex, Spin, Typography } from "antd";
import { InputField } from "../../components/form/InputField";

export const SigninPage = () => {

  const navigate = useNavigate({ from: "/auth/signin" });
  const signin = useMutation({
    mutationKey: ["signin"],
    mutationFn: authNetwork.signin,
    onSuccess: () => navigate({ to: "/" }),
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
      <p style={{ textAlign: "center" }}><Typography.Title>Sign In</Typography.Title></p>
      <Card style={{ maxWidth: "50vw", margin: "0 auto" }}>
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
              // validatorAdapter={zodValidator}
              // validators={{
              //   onBlur: signinInputSchema.shape.email,
              // }}
              children={(field) => <InputField field={field} label="Username:" />}
            />
            <form.Field
              name="passw"
              // validatorAdapter={zodValidator}
              // validators={{
              //   onBlur: signinInputSchema.shape.passw,
              // }}
              children={(field) => <InputField field={field} label="Password:" type="password" />}
            />
            {/* <form.Subscribe */}
            {/*   selector={(state) => state.errors} */}
            {/*   children={(errors) => errors.length > 0 && errors.toString()} */}
            {/* /> */}

            <br />
            <Button
              type="primary"
              htmlType="submit"
              disabled={signin.isPending}
            >
              {signin.isPending ? <Spin /> : "Submit"}
            </Button>
          </Flex>
        </form>
      </Card>
    </>
  );
}