import type { Route } from "./+types/overview";
import { fetchOverviewData } from "./utils/fetch-overivew-data";
import { Metric } from "./components/metric";

export const clientLoader = async () => {
  const overviewData = await fetchOverviewData();

  return { overviewData };
};

interface Props extends Route.ComponentProps {}

export default function Overview({ loaderData }: Props) {
  return (
    <article className="flex w-full h-full items-center justify-center gap-5">
      {loaderData.overviewData.metrics.map((metric, i) => (
        <Metric key={i} metric={metric} />
      ))}
    </article>
  );
}
