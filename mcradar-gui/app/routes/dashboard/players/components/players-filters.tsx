import { Input, Surface } from "@heroui/react";
import type { PlayersFiltersType } from "../types/players-filters-type";

export interface Props {
  setFilters: React.Dispatch<React.SetStateAction<PlayersFiltersType>>;
}

export const PlayersFilters = ({ setFilters }: Props) => {
  return (
    <Surface className="flex gap-2 p-2 rounded-lg">
      <Input
        type="text"
        variant="secondary"
        placeholder="Username"
        onChange={(e) =>
          setFilters((prev) => ({ ...prev, username: e.target.value }))
        }
      />
    </Surface>
  );
};
