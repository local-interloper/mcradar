import type { Column } from "~/shared/types/column";
import {
  FormatDate,
  FormatIp,
  FormatPlayerID,
  FormatPlayerName,
} from "~/shared/utils/formatters";

export const PLAYERS_COLUMNS: Column<Player>[] = [
  { label: "ID", property: "id", Formatter: FormatPlayerID },
  { label: "Name", property: "name", Formatter: FormatPlayerName },
  { label: "Created At", property: "createdAt", Formatter: FormatDate },
  { label: "Updated At", property: "updatedAt", Formatter: FormatDate },
];
