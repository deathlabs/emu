# Devices

Use `emu get devices` to get devices associated with the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Devices

```bash
emu get devices --system-id 1984
```

The response contains device identifiers and descriptive fields. When available, it also contains the date of the last scan.

```json
[
  {
    "deviceId": "ccf29d4c-6059-f111-b8f1-0050563303b9",
    "fqdn": "Application Layer Gateway",
    "hostname": "Application Layer Gateway",
    "lastScanDate": 1779840557.14,
    "nickname": "Application Layer Gateway"
  }
]
```

## Paginate Devices

Use `--page-index` to select the result page and `--page-size` to set the maximum number of devices in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get devices \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```
