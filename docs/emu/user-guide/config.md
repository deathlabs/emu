# Configuration Output

Use `emu get config` to print the configuration loaded by EMU. This command reads local configuration and does not send a GET request to eMASS.

```bash
emu get config
```

The output contains the configured `url`, `profiles`, `systems`, and `settings`. Review the output before you share it because it can contain local certificate paths, profile names, and system information.

Use `--output` to select JSON or YAML output:

```bash
emu get config --output json
```
