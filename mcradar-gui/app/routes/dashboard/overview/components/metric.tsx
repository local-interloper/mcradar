import { Card } from "@heroui/react";
import type { MetricData } from "../types/metric-data";

interface Props {
  metric: MetricData;
}

export const Metric = ({ metric: { label, value, Icon } }: Props) => {
  return (
    <Card variant="tertiary">
      <Card.Content>
        <article className="relative flex flex-col items-center justify-center w-16 h-16">
          <span className="absolute z-10 text-4xl font-semibold">{value}</span>
          {Icon && <Icon className="absolute text-7xl opacity-10" />}
        </article>
      </Card.Content>
      <Card.Footer>
        <span className="font-semibold text-lg">{label}</span>
      </Card.Footer>
    </Card>
  );
};
