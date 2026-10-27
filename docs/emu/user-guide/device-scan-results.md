# Device Scan Results

Use `emu get device-scan-results` to get device scan results associated with the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Device Scan Results

```bash
emu get device-scan-results --system-id 1984
```

Each result can contain the following fields:

* `benchmark`
* `benchmarkRelease`
* `deviceId`
* `firstSeenDate`
* `ingestedDate`
* `lastScanDate`
* `scanResult`
* `scanType`
* `securityCheck`

## Paginate Device Scan Results

Use `--page-index` to select the result page and `--page-size` to set the maximum number of scan results in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get device-scan-results \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```
