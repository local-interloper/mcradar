import { API_URL } from "~/config";
import { getKey } from "./auth-utils";
import type { TableDataRequest } from "../types/table-data-request";

export const fetchWithKey = async (endpoint: string, init?: RequestInit) => {
  return fetch(`${API_URL}${endpoint}`, {
    ...init,
    headers: {
      ...init?.headers,
      Authorization: `Bearer ${getKey()}`,
    },
  });
};

export const fetchTableData = async <T>({
  endpoint,
  pagination,
  filters,
}: TableDataRequest) => {
  const response = await fetchWithKey(endpoint, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      pagination,
      filters,
    }),
  });

  return (await response.json()) as TableData<T>;
};
