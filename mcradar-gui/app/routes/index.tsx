import { Spinner } from "@heroui/react";
import type { Route } from "./+types/index";
import { useEffect } from "react";
import { isKeyValid } from "~/shared/utils/auth-utils";
import { redirect, useNavigate } from "react-router";

interface Props extends Route.ComponentProps {}

export const clientLoader = async () => {
  if (await isKeyValid()) {
    return redirect("/dashboard/overview");
  }

  return redirect("/login");
};

const Index = ({}: Props) => {
  return (
    <article className="flex w-full h-full items-center justify-center">
      <Spinner size="xl" />
    </article>
  );
};

export default Index;
