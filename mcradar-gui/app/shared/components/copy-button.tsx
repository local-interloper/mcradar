import { Button, Tooltip } from "@heroui/react";
import { HiOutlineClipboardCopy } from "react-icons/hi";

interface Props {
  text: string;
}

export const CopyButton = ({ text }: Props) => {
  const copyHandler = () => {
    navigator.clipboard.writeText(text);
  };

  return (
    <Tooltip delay={0} closeDelay={200}>
      <Button size="sm" variant="ghost" onClick={copyHandler}>
        <HiOutlineClipboardCopy />
      </Button>
      <Tooltip.Content>
        <span>Copy</span>
      </Tooltip.Content>
    </Tooltip>
  );
};
