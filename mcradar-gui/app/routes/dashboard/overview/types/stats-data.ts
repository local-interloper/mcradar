import type { MetricData } from "./metric-data";

export interface StatsData {
    label: string;
    metrics: MetricData[];
}