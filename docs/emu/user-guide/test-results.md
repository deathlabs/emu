
# Test Results

## Get Test Results

Use `emu get test-results` to get control test results for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

```bash
emu get test-results --system-id 1984
```

Each result can contain fields such as `control`, `assessmentProcedure`, `cci`, `complianceStatus`, `description`, `testDate`, and `testedBy`.

## Filter Test Results

Use the test result filters separately or together:

* `--control-acronyms`: Filter by one or more control acronyms.
* `--assessment-procedures`: Filter by one or more assessment procedures.
* `--ccis`: Filter by one or more Control Correlation Identifiers (CCIs).
* `--latest-only`: Return only the latest test result for each control.

```bash
emu get test-results \
  --system-id 1984 \
  --control-acronyms AC-1 \
  --latest-only
```

## Upload a Test Result
```bash
emu upload test-results \
  -compliance-status Compliant \
  -assessment-procedure "AC-1.1" \
  -description "Testing" \
  -tested-by "Fernandez, Victor" \
  -test-date "2024-08-22 15:06:06"
```
