import { useState } from "react";
import { TbTrafficCone } from "react-icons/tb";
import type { PlayersFiltersType } from "./types/players-filters-type";
import { DataTable } from "~/shared/components/data-table";
import { PLAYERS_COLUMNS } from "./config/players-columns";
import { PlayersFilters } from "./components/players-filters";

const Players = () => {
  const [filters, setFilters] = useState<PlayersFiltersType>({});

  return (
    <article className="flex gap-2 flex-col w-full h-full">
      <PlayersFilters setFilters={setFilters} />
      <DataTable
        endpoint="/players"
        columns={PLAYERS_COLUMNS}
        filters={filters}
      />
    </article>
  );
};

export default Players;
