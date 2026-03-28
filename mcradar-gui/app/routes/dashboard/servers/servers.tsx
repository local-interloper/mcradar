import { SERVERS_COLUMNS } from "./config/servers-columns";
import { Suspense, useState } from "react";
import type { ServersFiltersType } from "./types/servers-filters-type";
import { DataTable } from "~/shared/components/data-table";
import { ServersFilters } from "./components/servers-filters";
import { Spinner } from "@heroui/react";

const Servers = () => {
  const [filters, setFilters] = useState<ServersFiltersType>({});

  return (
    <article className="flex gap-2 flex-col w-full h-full">
      <ServersFilters setFilters={setFilters} />

      <Suspense fallback={<Spinner />}>
        <DataTable
          endpoint="/servers"
          columns={SERVERS_COLUMNS}
          filters={filters}
        />
      </Suspense>
    </article>
  );
};

export default Servers;
