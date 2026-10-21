# Container Scan Results

## Upload a Container Scan Result 
```bash
emu upload container-scan-results \
  -container-id 157ac21c-49b4-4faf-a4ac-bfefd869ba3b \
  -container-name "aide-engine" \
  -time "2024-09-10 14:00:01" \
  -benchmark "RHEL_8_STIG" \
  -filepath demo.csv
```

## Delete a Container Scan Result 
```bash
emu delete container-scan-results \
  -container-id 157ac21c-49b4-4faf-a4ac-bfefd869ba3b
```
