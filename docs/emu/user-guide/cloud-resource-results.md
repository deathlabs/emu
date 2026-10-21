# Cloud Resource Results

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
