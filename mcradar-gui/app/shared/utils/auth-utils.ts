import { API_URL } from "~/config";

const API_KEY_LOCAL_STORAGE_KEY = "API_KEY";

export const setKey = (key: string) => {
  localStorage.setItem(API_KEY_LOCAL_STORAGE_KEY, key);
};

export const getKey = (): string => {
  return localStorage.getItem(API_KEY_LOCAL_STORAGE_KEY) || "";
};

export const clearKey = () => {
  localStorage.removeItem(API_KEY_LOCAL_STORAGE_KEY);
}

export const isKeyValid = async (): Promise<boolean> => {
  const response = await fetch(`${API_URL}/auth`, {
    headers: {
      "Authorization": `Bearer ${getKey()}`
    }
  });
  return response.ok;
};
