import { Button, Input, Surface, toast } from "@heroui/react";
import type { Route } from "./+types/index";
import { useNavigate } from "react-router";
import { isKeyValid, setKey } from "~/shared/utils/auth-utils";
import { useState } from "react";

interface Props extends Route.ComponentProps {}

const Login = ({}: Props) => {
  const navigate = useNavigate();
  const [keyInput, setKeyInput] = useState("");

  const loginHandler = async () => {
    setKey(keyInput);

    if (await isKeyValid()) {
      navigate("/dashboard/overview");
      return;
    }

    toast("Invalid API key", {
      variant: "danger",
    });
  };

  return (
    <article className="flex w-full h-full items-center justify-center">
      <Surface className="class flex flex-col p-2 gap-4 rounded-lg items-center">
        <h1 className="font-bold text-2xl">Login</h1>
        <Input
          variant="secondary"
          type="password"
          placeholder="API Key"
          onChange={(e) => setKeyInput(e.target.value)}
        />
        <Button onClick={loginHandler}>Login</Button>
      </Surface>
    </article>
  );
};

export default Login;
