# Container Scan Results

## Get Container Scan Results

Use `emu get container-scan-results` to get container scan results for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

```bash
emu get container-scan-results --system-id 1984
```

The command returns an empty list when the system has no container scan results:

```json
[]
```

## Paginate Container Scan Results

Use `--page-index` to select the result page and `--page-size` to set the maximum number of results in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get container-scan-results \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```

## Upload a Container Scan Result 
```bash
emu upload container-scan-results \
  -container-id 157ac21c-49b4-4faf-a4ac-bfefd869ba3b \
  -container-name "aide-engine" \
  -time "2024-09-10 14:00:01" \
  -benchmark "RHEL_8_STIG" \
  -filepath demo.csv
```

## Delete a Container Scan Result 
```bash
emu delete container-scan-results \
  -container-id 157ac21c-49b4-4faf-a4ac-bfefd869ba3b
```
