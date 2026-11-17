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

## Update Artifacts

The update operation replaces all artifact metadata. Include all values that you want to keep.

```bash
emu update artifact \
  --system-id 1984 \
  --file artifact-update.yaml
```

The input file must contain a list of artifact records:

```yaml
- filename: policy.pdf
  name: Security Policy
  description: Approved security policy
  isTemplate: false
  type: Policy
  category: Evidence
  referencePageNumber: Page 1
  controls: AC-1,AC-2
  assessmentProcedures: AC-1.1,AC-1.2
  expirationDate: 1790812800
  lastReviewedDate: 1759276800
  signedDate: 1759276800
```

## Delete Artifacts

```bash
emu delete artifact \
  --system-id 1984 \
  --file artifact-delete.yaml
```

The input file must contain a list of artifact filenames:

```yaml
- filename: policy.pdf
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
