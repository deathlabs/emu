# Cloud Resource Results

## Get Cloud Resource Results

Use `emu get cloud-resource-results` to get cloud resource compliance results for the configured systems. Use the global `--system-id` flag to limit the request to one or more systems.

```bash
emu get cloud-resource-results --system-id 1984
```

The command returns an empty list when the system has no cloud resource results:

```json
[]
```

## Paginate Cloud Resource Results

Use `--page-index` to select the result page and `--page-size` to set the maximum number of results in the page. The default page index is `0`, and the default page size is `20000`.

```bash
emu get cloud-resource-results \
  --system-id 1984 \
  --page-index 0 \
  --page-size 5
```

## Upload a Cloud Resource Result
```bash
emu upload cloud-resource-results \
  -provider=azure \
  -resource-id=/subscriptions/123456789/sample/resource/namespace/default \
  -resource-name=Storage Resource \
  -resource-type=Microsoft.storage.table \
  -initiated-by=john.doe.ctr@mail.mil \
  -csp-account-id=123456789 \
  -csp-region=useast2 \
  -tags=test:testtag \
  -csp-policy-definition-id="/providers/sample/policy/namespace/au11_policy" \
  -policy-definition-title="AU-11 - Audit Record Retention" \
  -compliance-check-timestamp=1644003780 \
  -is-compliant=true \
  -control="AU-11" \
  -assessment-procedure=000167,000168 \
  -policy-deployment-name=testDeployment \
  -policy-deployment-version=1.0.0 \
  -severity=Low
```

## Delete a Cloud Resource Result  
```bash
emu delete cloud-resource-results \
  -resource-id "/subscriptions/123456789/sample/resource/namespace/default"
```
