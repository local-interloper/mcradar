import type { ServerType } from "~/shared/types/server-type";

export interface ServersFiltersType {
  version?: string;
  types?: ServerType[];
}
