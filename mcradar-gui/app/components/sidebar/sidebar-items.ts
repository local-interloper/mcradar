import type { ReactElement, ReactNode } from "react";
import type { IconType } from "react-icons";
import { HiEye, HiServer, HiUser } from "react-icons/hi2";

export interface SidebarItem {
  label: string;
  Icon: IconType;
  url: string;
}

export const SIDEBAR_ITEMS: SidebarItem[] = [
  {
    label: "Overview",
    Icon: HiEye,
    url: "/dashboard/overview",
  },
  {
    label: "Servers",
    Icon: HiServer,
    url: "/dashboard/servers",
  },
  {
    label: "Players",
    Icon: HiUser,
    url: "/dashboard/players",
  },
];
