import { Spinner } from "@heroui/react";
import { Suspense } from "react";
import { Outlet, redirect } from "react-router";
import { Sidebar } from "~/components/sidebar/sidebar";
import { isKeyValid } from "~/shared/utils/auth-utils";

export const clientLoader = async () => {
  if (!(await isKeyValid())) {
    return redirect("/login");
  }
};

const DashboardLayout = () => {
  return (
    <article className="flex w-full h-full">
      <Sidebar />
      <main className="flex flex-col w-full h-full p-2 overflow-hidden ">
        <Outlet />
      </main>
    </article>
  );
};

export default DashboardLayout;
