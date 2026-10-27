# Applications

Use `emu get applications` to get the applications associated with the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

## Get Applications

```bash
emu get applications --system-id 1984
```

The command returns an empty list when the system has no applications:

```json
[]
```

## Paginate Applications

Use `--page-index` to select the result page and `--page-size` to set the maximum number of applications in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get applications \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```
