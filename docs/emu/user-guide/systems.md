# Systems

Use `emu get systems` to get systems available to the selected EMU profiles. Use the global `--system-id` flag to request a specific configured system.

## Get Systems

```bash
emu get systems --profile production
```

The command returns a list when you do not specify a system ID.

## Get a System

```bash
emu get systems --system-id 1984
```

When you specify one system ID, the response is a single object that includes fields such as `name`, `acronym`, `systemId`, `policy`, `authorizationStatus`, and `registrationType`.

```json
{
  "acronym": "Example",
  "authorizationStatus": "Authorization to Operate (ATO)",
  "name": "Example System",
  "policy": "RMF",
  "registrationType": "Assess and Authorize",
  "systemId": 1984
}
```

## Filter Systems

The command supports these filters:

* `--coams-id`: Return systems with a COAMS ID.
* `--ditpr-id`: Return systems with a DITPR ID.
* `--include-decommissioned`: Include decommissioned systems.
* `--include-ditpr-metrics`: Include DITPR metrics. Do not combine this flag with `--coams-id` or `--ditpr-id`.
* `--policy`: Filter by `DIACAP`, `RMF`, or `Reporting`.
* `--registration-type`: Filter by an eMASS registration type.
* `--reports-for-scorecard`: Return only systems that report to the DoD Cyber Hygiene Scorecard.

```bash
emu get systems \
  --profile production \
  --policy RMF \
  --reports-for-scorecard
```
