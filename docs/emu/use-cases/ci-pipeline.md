# Updating eMASS in a CI Pipeline

```yaml
---
stages:
  - deliver

push-emass-artifact:
  stage: deliver
  image: registry.cdso.army.mil/cdso/containers/approved-tool/emu:latest
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH
  variables:
    EMASS_API_URL: "https://connect.emass.apps.mil"
    EMASS_API_UID: $EMASS_API_UID
    EMASS_API_KEY: $EMASS_API_KEY
    EMASS_SYSTEM_ID: $EMASS_SYSTEM_ID
    NPE_CERTIFICATE: $NPE_CERTIFICATE
    NPE_KEY: $NPE_KEY
  script:
    - |
      echo "[*] Updating eMASS system record ${EMASS_SYSTEM_ID}"
      emu upload artifact -filename policy.pdf
tags:
  - cdso
```
