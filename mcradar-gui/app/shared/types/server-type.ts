export const SERVER_TYPES = ["legit", "cracked", "unknown"] as const;

export type ServerType = typeof SERVER_TYPES[number];
