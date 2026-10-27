# Artifacts

## Upload an Artifact
```bash
emu upload artifact \
  -filename policy.pdf
```

## Delete an Artifact
```bash
emu delete artifact \
  -filename policy.pdf
```

## Upload Multiple Files
```bash
emu upload artifacts \
  -filename policies.zip -is-bulk=true
```

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
