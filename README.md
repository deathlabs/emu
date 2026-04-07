# `emu`
A tool for automating eMASS records management.

## Quickstart
**Step 1.** Compile the source code and install `emu`. 
```bash
make
```

**Step 2.** Set your eMASS API keys as environment variables. 
```bash
# Profile 1
export EMASS_USER_UID_PRODUCTION="1234567890"
export EMASS_API_KEY_PRODUCTION="aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"

# Profile 2
export EMASS_USER_UID_PILOT="0987654321"
export EMASS_API_KEY_PILOT="bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
```

**Step 3.** Create an `emu` configuration file. Below is an example of what it should look like. 
```yaml
---
url: http://localhost
profiles:
  - name: production
    publicKeyPath: production.cer
    privateKeyPath: production.key
  - name: pilot
    publicKeyPath: pilot.cer
    privateKeyPath: pilot.key
systems:
  - name: Skynet
    id: 1984
    profile: production
  - name: Genisys
    id: 2015
    profile: pilot
  - name: Legion
    id: 2019
    profile: pilot
settings:
  output:
    format: yaml
```

**Step 4.** Run `emu`.
```bash
emu get systems
```

You should get output similar to below. 
```json
[
  {
    "acronym": "SKY",
    "description": "Autonomous defense network managed by Cyberdyne Systems.",
    "name": "Skynet",
    "owningOrganization": "Cyberdyne Systems",
    "policy": "RMF",
    "registrationType": "Assess and Authorize",
    "securityPlanApprovalDate": 1638741660,
    "securityPlanApprovalStatus": "Approved",
    "systemId": 1984
  },
  {
    "acronym": "LGN",
    "description": "Next-generation autonomous threat platform managed by Cyberdyne Systems.",
    "name": "Legion",
    "owningOrganization": "Cyberdyne Systems",
    "policy": "RMF",
    "registrationType": "Assess and Authorize",
    "securityPlanApprovalDate": null,
    "securityPlanApprovalStatus": "Not Yet Approved",
    "systemId": 2019
  },
  {
    "acronym": "GNS",
    "description": "Integrated global operating system managed by Cyberdyne Systems.",
    "name": "Genisys",
    "owningOrganization": "Cyberdyne Systems",
    "policy": "RMF",
    "registrationType": "Assess and Authorize",
    "securityPlanApprovalDate": 1700000000,
    "securityPlanApprovalStatus": "Approved",
    "systemId": 2015
  }
]
```