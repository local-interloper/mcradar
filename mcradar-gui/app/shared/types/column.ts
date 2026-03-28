import type { FC } from "react";
import type { FormatterProps } from "../utils/formatters";

export interface Column<T> {
  label: string;
  property: keyof T;
  Formatter?: FC<FormatterProps>;
}
