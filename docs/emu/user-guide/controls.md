
# Controls

Use `emu get controls` to get controls for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Controls

```bash
emu get controls --system-id 1984
```

The command returns an empty list when no controls match the request:

```json
[]
```

## Filter Controls

Use `--control-acronyms` to request one or more controls.

```bash
emu get controls \
  --system-id 1984 \
  --control-acronyms SI-12,AC-1,TR-1
```
