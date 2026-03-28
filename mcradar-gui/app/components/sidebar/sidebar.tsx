import { Button, Surface, Tooltip } from "@heroui/react";
import SidebarButton from "./sidebar-button";
import { SIDEBAR_ITEMS } from "./sidebar-items";
import { clearKey } from "~/shared/utils/auth-utils";
import { useNavigate } from "react-router";
import { HiPower } from "react-icons/hi2";

export const Sidebar = () => {
  const navigate = useNavigate();

  const logoutHandler = () => {
    clearKey();
    navigate("/");
  };

  return (
    <Surface className="flex flex-col w-1/6 gap-5 p-2 pt-5">
      <section className="flex justify-center">
        <section className="relative flex flex-col items-center w-60 h-60">
          <img src="/logo.png" alt="mcradar logo" className="opacity-60" />
          <h1 className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-3xl font-bold">
            mcradar
          </h1>
        </section>
      </section>
      <section className="flex flex-col gap-2 h-full w-full">
        {SIDEBAR_ITEMS.map((item, i) => (
          <SidebarButton key={i} item={item} />
        ))}
      </section>
      <section className="flex">
        <Tooltip delay={0}>
          <Button
            className="w-10 h-10"
            onClick={logoutHandler}
            variant="danger-soft"
          >
            <HiPower />
          </Button>
          <Tooltip.Content placement="right">
            <p>Log Out</p>
          </Tooltip.Content>
        </Tooltip>
      </section>
    </Surface>
  );
};
