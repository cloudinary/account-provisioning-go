# MetricUsage


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Metric`                                                                     | `*string`                                                                    | :heavy_minus_sign:                                                           | The type of metric being measured (e.g., "credit").                          |
| `Limit`                                                                      | `*float64`                                                                   | :heavy_minus_sign:                                                           | The limit for this metric in the plan.                                       |
| `Cycles`                                                                     | [][components.Cycles](../../models/components/cycles.md)                     | :heavy_minus_sign:                                                           | Full usage details per billing cycle (monthly or annual, depending on plan). |
| `MonthlyCycles`                                                              | [][components.MonthlyCycles](../../models/components/monthlycycles.md)       | :heavy_minus_sign:                                                           | Monthly-level breakdown within a single annual cycle.                        |