type History = Array<{
  status: 200;
  expected: 200;
  timestamp: string;
  duration: 1078180125;
}>;

type Stats = {
  total_checks: 48;
  successful_checks: 48;
  uptime_percentage: 100;
  average_response_ms: 237;
  last_check: string;
};

export type Endpoint = {
  url: string;
  history: History;
  stats: Stats;
};

export type Endpoints = {
  domain: string;
  endpoints: Array<Endpoint>;
};
