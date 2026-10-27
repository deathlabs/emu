# System Roles

Use `emu get system-roles` to get the roles that eMASS supports for the Package Approval Chain (PAC), Control Approval Chain (CAC), and other assignments. This request applies to an EMU profile rather than a specific system.

## Get System Roles

```bash
emu get system-roles --profile production
```

The response contains each role and its category. Example values include `AO`, `AODR`, `Auditor`, `ISO/PM`, `Validate`, and `View Only`.

```json
[
  {
    "role": "AO",
    "roleCategory": "PAC"
  },
  {
    "role": "Auditor",
    "roleCategory": "Other"
  }
]
```

## Filter System Roles

Use `--category` with `--role` to request a specific role. Valid category examples are `PAC`, `CAC`, and `Other`. You can also use `--policy` to specify `RMF`, `DIACAP`, or `Reporting`.

```bash
emu get system-roles \
  --profile production \
  --category PAC \
  --role AO \
  --policy RMF
```

The `--role` flag is required when you use `--category`.
