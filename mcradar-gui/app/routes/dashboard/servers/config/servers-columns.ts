import type { Column } from "~/shared/types/column";
import {
  FormatDate,
  FormatIp,
  FormatServerType,
} from "~/shared/utils/formatters";

export const SERVERS_COLUMNS: Column<Server>[] = [
  { label: "IP", property: "ip", Formatter: FormatIp },
  { label: "Version", property: "version" },
  { label: "Type", property: "type", Formatter: FormatServerType },
  { label: "Online Players", property: "onlinePlayers" },
  { label: "Max Players", property: "maxPlayers" },
  { label: "Created At", property: "createdAt", Formatter: FormatDate },
  { label: "Updated At", property: "updatedAt", Formatter: FormatDate },
];
