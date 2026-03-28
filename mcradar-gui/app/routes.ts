import {
  type RouteConfig,
  index,
  layout,
  prefix,
  route,
} from "@react-router/dev/routes";

export default [
  layout("layouts/core-layout.tsx", [
    index("routes/index.tsx"),
    route("login", "routes/login.tsx"),
    layout("layouts/dashboard-layout.tsx", [
      ...prefix("dashboard", [
        route("overview", "routes/dashboard/overview/overview.tsx"),
        route("servers", "routes/dashboard/servers/servers.tsx"),
        route("players", "routes/dashboard/players/players.tsx"),
      ]),
    ]),
  ]),
] satisfies RouteConfig;
