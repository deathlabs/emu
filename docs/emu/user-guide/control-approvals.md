# CAC Control Approvals

Use `emu get control-approvals` to get control approvals in the Control Approval Chain (CAC) for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Control Approvals

```bash
emu get control-approvals --system-id 1984
```

The command returns an empty list when the system has no matching control approvals:

```json
[]
```

## Filter Control Approvals

Use `--control-acronyms` to request approvals for one or more controls.

```bash
emu get control-approvals \
  --system-id 1984 \
  --control-acronyms AC-1,AC-2
```
