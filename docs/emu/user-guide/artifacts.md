# Artifacts

## Upload an Artifact
```bash
emu upload artifact \
  --system-id 1984 \
  --file policy.pdf
```

## Upload Multiple Artifacts

To upload multiple artifacts, put the files in a ZIP archive and use `--is-bulk`.

```bash
emu upload artifact \
  --system-id 1984 \
  --file policies.zip \
  --is-bulk
```

## Update and Delete Limitations

The `emu update artifact` and `emu delete artifact` commands are placeholders. They do not send requests to eMASS, and they cannot update or delete artifacts in the current release.

## Get Artifacts

Use `emu get artifacts` to get artifact metadata for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

```bash
emu get artifacts --system-id 1984
```

Each result can contain fields such as `name`, `filename`, `category`, `description`, `controls`, `assessmentProcedures`, `fileSize`, and `mimeContentType`.

## Filter Artifacts

Use the artifact filters separately or together:

* `--filename`: Filter by filename.
* `--control-acronyms`: Filter by one or more control acronyms.
* `--assessment-procedures`: Filter by one or more assessment procedures.
* `--ccis`: Filter by one or more Control Correlation Identifiers (CCIs).
* `--system-only`: Exclude control-level and assessment-procedure-level artifacts.

```bash
emu get artifacts \
  --system-id 1984 \
  --filename policy.pdf \
  --system-only
```
