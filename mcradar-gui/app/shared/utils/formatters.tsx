import { Button, Chip, toast, Tooltip, TooltipContent } from "@heroui/react";
import type { ReactNode } from "react";
import { HiCheck } from "react-icons/hi2";
import { CopyButton } from "../components/copy-button";

export interface FormatterProps {
  children: string;
}

export const FormatDate = ({ children }: FormatterProps): ReactNode => {
  const nativeDate = new Date(children);

  const [year, month, day, hours, minutes] = [
    nativeDate.getFullYear(),
    String(nativeDate.getMonth() + 1).padStart(2, "0"),
    String(nativeDate.getDate()).padStart(2, "0"),
    String(nativeDate.getHours()).padStart(2, "0"),
    String(nativeDate.getMinutes()).padStart(2, "0"),
  ];

  const date = `${year}-${month}-${day}`;
  const time = `${hours}:${minutes}`;

  return (
    <article className="flex items-center gap-2">
      <span>{date}</span>
      <span>{time}</span>
    </article>
  );
};

export const FormatServerType = ({ children }: FormatterProps): ReactNode => {
  switch (children) {
    case "legit":
      return (
        <Chip className="w-20 justify-center" color="success">
          Legit
        </Chip>
      );
    case "cracked":
      return (
        <Chip className="w-20 justify-center" color="warning">
          Cracked
        </Chip>
      );
    default:
      return (
        <Chip className="w-20 justify-center" color="danger">
          Unknown
        </Chip>
      );
  }
};

export const FormatIp = ({ children }: FormatterProps): ReactNode => {
  const onCopyHandler = () => {
    navigator.clipboard.writeText(children);
    toast("IP copied to clipboard!", {
      variant: "success",
      indicator: <HiCheck />,
    });
  };

  return (
    <article className="flex items-center gap-2">
      <CopyButton text={children} />
      <span className="font-mono opacity-70">{children}</span>
    </article>
  );
};

export const FormatPlayerID = ({ children }: FormatterProps): ReactNode => {
  const startEndLength = 5;

  const copyHandler = () => {
    navigator.clipboard.writeText(children);
    toast("Player ID copied to clipboard!", {
      variant: "success",
      indicator: <HiCheck />,
    });
  };

  const abbreviated = `${children.substring(0, startEndLength)}...${children.substring(children.length - startEndLength)}`;

  return (
    <article className="flex items-center gap-2">
      <CopyButton text={children} />
      <Tooltip delay={0} closeDelay={0}>
        <Tooltip.Trigger>
          <span className="font-mono opacity-70">{abbreviated}</span>
        </Tooltip.Trigger>
        <Tooltip.Content showArrow>
          <span className="font-mono">{children}</span>
        </Tooltip.Content>
      </Tooltip>
    </article>
  );
};

export const FormatPlayerName = ({ children }: FormatterProps): ReactNode => {
  const sanitized = children.split(/§./).join("");

  return <span>{sanitized}</span>;
};
