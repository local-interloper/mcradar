import { fetchWithKey } from "~/shared/utils/fetching-utils";
import type { OverviewResponse } from "../types/overview-response";
import type { OverviewData } from "../types/overview-data";
import { HiServer, HiUser } from "react-icons/hi2";

export const fetchOverviewData = async (): Promise<OverviewData> => {
  const response = await fetchWithKey("/overview");
  const raw = (await response.json()) as OverviewResponse;

  return {
    metrics: [
      {
        label: "Servers",
        value: raw.totalServers,
        Icon: HiServer,
      },
      {
        label: "Players",
        value: raw.totalPlayers,
        Icon: HiUser,
      },
    ],
  };
};
