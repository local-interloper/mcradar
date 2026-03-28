import { cn } from "@heroui/styles";
import type { SidebarItem } from "./sidebar-items";
import { NavLink, useLocation } from "react-router";

interface Props {
  item: SidebarItem;
}

const SidebarButton = ({ item: { label, Icon, url } }: Props) => {
  const { pathname } = useLocation();

  const isActive = pathname === url;

  return (
    <NavLink
      className={cn(
        "button gap-2 justify-start w-full",
        isActive ? "button--primary" : "button--ghost",
      )}
      to={url}
    >
      <Icon />
      <span>{label}</span>
    </NavLink>
  );
};

export default SidebarButton;
