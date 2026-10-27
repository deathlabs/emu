
# Static Code Scans

## Get Static Code Scans

Use `emu get static-code-scans` to get static code scans for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

```bash
emu get static-code-scans --system-id 1984
```

The command returns an empty list when the system has no static code scans:

```json
[]
```

## Paginate Static Code Scans

Use `--page-index` to select the result page and `--page-size` to set the maximum number of scans in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get static-code-scans \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```

## Upload a Static Code Scan
```bash
emu upload static-code-scans upload \
  -applicationName="aide" \
  -clearFindings=false \
  -codeCheckName="Aide" \
  -count="12" \
  -cweId="155" \
  -rawSeverity="low" \
  -scanDate="2024-09-18 11:05:30" \
  -version="9002"
```
