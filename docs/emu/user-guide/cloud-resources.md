# Cloud Resources

Use `emu get cloud-resources` to get cloud resources associated with the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Cloud Resources

```bash
emu get cloud-resources --system-id 1984
```

The command returns an empty list when the system has no cloud resources:

```json
[]
```

## Paginate Cloud Resources

Use `--page-index` to select the result page and `--page-size` to set the maximum number of cloud resources in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get cloud-resources \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```
