# Device Scans
The table below describes what file types are currently supported by eMASS and `emu`.

## Supported Scan Types

| Scan Type               | File Type | Supported by eMASS | Supported by EMU |
| ----------------------- | --------- | ------------------ | ---------------- |
| `disaStigViewerCklCklb` | `.cklb`   | Yes                | Yes              |
| `acasNessus`            | `.nessus` | Yes                | Yes              |
| `disaStigViewerCmrs`    | `.cmrs`   | Yes                | TBD              |
| `acasAsrArf`            | `.zip`    | Yes                | Yes              |
| `acasConslidatedArf`    | `.arf`    | Yes                | TBD              |
| `acasNessus`            | `.csv`    | No                 | TBD              |
| `acasNessus`            | `.html`   | No                 | TBD              |
| `acasNessus`            | `.pdf`    | No                 | TBD              |
| `scapComplianceChecker` | TBD       | Yes                | TBD              |
| `policyAuditor`         | TBD       | Yes                | TBD              |

## Uploading a STIG Checklist
```bash
emu upload device-scans \
  -filename scans/device-scan.cklb \
  -scan-type disaStigViewerCklCklb
```

## Uploading an ACAS Scan  
```bash
emu upload device-scans \
  -filename scans.nessus \
  -scan-type acasNessus
```

```bash
emu upload device-scan \
  -filename results.zip \
  -scan-type acasAsrArf
```
