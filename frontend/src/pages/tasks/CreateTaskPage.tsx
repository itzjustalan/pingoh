import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { zodValidator } from "@tanstack/zod-form-adapter";
import { Button, Card, Divider, Flex, Spin, Typography, message } from "antd";
import { InputField } from "../../components/form/InputField";
import { createTaskSchema } from "../../lib/models/db/task";
import { tasksNetwork } from "../../lib/networks/tasks";

const defaultValues = {
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
};

export const CreateTaskPage = () => {
  const [messageApi, contextHolder] = message.useMessage();
  const createTask = useMutation({
    mutationKey: ["create", "task"],
    mutationFn: tasksNetwork.create,
    onSuccess: () => {
      messageApi.success("Task created");
      form.reset();
    },
    onError: (error) => {
      messageApi.error(error.message ?? "Error creating Task");
    },
  });
  const form = useForm({
    defaultValues,
    validatorAdapter: zodValidator(),
    validators: {
      onChange: createTaskSchema,
    },
    onSubmit: async ({ value }) => {
      console.log(value);
      // messageApi.warning("Validation failed")
      createTask.mutateAsync(value);
    },
  });

  // const fillForm = () => {
  //   form.setFieldValue("name", "test");
  //   form.setFieldValue("interval", 60);
  //   form.setFieldValue("description", "test task");
  //   form.setFieldValue("http.url", "https://example.com");
  // };

  return (
    <>
      {contextHolder}
      <Typography.Title level={2}>Create Task</Typography.Title>
      <Divider />
      {/* <Button htmlType="button" onClick={fillForm}> */}
      {/*   Fill */}
      {/* </Button> */}
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
              name="name"
              children={(field) => <InputField field={field} label="Name:" />}
            />
            <br />
            <form.Field
              name="interval"
              children={(field) => (
                <InputField field={field} type="number" label="Interval:" />
              )}
            />
            <br />
            <form.Field
              name="description"
              children={(field) => (
                <InputField field={field} label="Description:" />
              )}
            />
            <br />
            <form.Field
              mode="array"
              name="http.url"
              children={(field) => (
                <InputField field={field} type="url" label="Url:" />
              )}
            />
            <br />

            <Button
              type="primary"
              htmlType="submit"
              disabled={createTask.isPending}
            >
              {createTask.isPending ? <Spin /> : "Submit"}
            </Button>
          </Flex>
        </form>
      </Card>
    </>
  );
};
