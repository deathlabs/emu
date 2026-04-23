# `emu`
The eMASS Updater (EMU) is a CLI tool for automating eMASS records management.

## Installation
EMU can be installed using one of two options. 

**Option 1.** Download the latest pre-built binary from [EMU's Releases page](https://github.com/deathlabs/emu/releases), make it executable, and then put it in your execution path (below is an Linux-based example showing how to make it executable and where to put it).
```bash
chmod +x emu
mv emu /usr/local/bin/
```

**Option 2.** Clone EMU's GitHub repository, change directories to the folder downloaded, and then, build EMU locally using the provided Makefile. Like Option 1, make sure to make EMU executable and put it somewhere in your execution path (the example below is also Linux-based but the idea is generally same for Windows). 
```bash
git clone https://github.com/deathlabs/emu.git
cd emu
make
chmod +x emu
mv emu /usr/local/bin/
```

## Configuration
An EMU configuration profile represents a unique group of values and preferences. You can use multiple profiles to filter which eMASS records you want to interact with, which POA&M strategies you want to use, etc. 

By default, EMU will use all profiles specified in your EMU configuration file and therefore interact with all eMASS records associated with them. 

Below is an example of what your EMU configuration file should look like if you're using the DOD's instance of eMASS, have two EMU configuration profiles called `production` and `pilot`, and have three eMASS records. If you're not using the DOD's instance of eMASS, change the URL to what it should be for the eMASS API. For each eMASS record, make sure to specify the EMU configuration profile you want to use with it. 
```yaml
---
url: https://connect.emass.apps.mil
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
    format: json
```

## Quickstart
**Step 1.** After installing EMU (see the [Installation](#installation) instructions above), set your eMASS User UIDs (e.g., DOD ID number) and eMASS API keys as environment variables. 

EMU expects your User UID to start with `EMASS_USER_UID_` and end with the name of your EMU configuration profile in uppercase (e.g., `PRODUCTION`, `PILOT`, `GROUP_1`, `CLOUD_SYSTEMS`, or `EDGE_SYSTEMS`). The EMU configuration profile names will only matter to you, the user. 

For your API key, the same rule applies. It must start with `EMASS_API_KEY_` and end with the name of your EMU configuration profile in uppercase. The example below shows how to set your environment variables if you were using a Linux-based environment and had two EMU configuration profiles called `PRODUCTION` and `PILOT`. 
```bash
# Profile 1
export EMASS_USER_UID_PRODUCTION="1234567890"
export EMASS_API_KEY_PRODUCTION="aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"

# Profile 2
export EMASS_USER_UID_PILOT="0987654321"
export EMASS_API_KEY_PILOT="bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
```

**Step 2.** Create an EMU configuration file if you haven't already (see the [Configuration](#configuration) instructions above). Make sure to either name your EMU configuration file `.emu.yaml` or specify the file path to it when you run EMU. 

**Step 3.** Run an EMU command.
```bash
emu get systems
```

You should get output similar to below depending on which command you ran.  
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

**Step 4.** If you wanted to run EMU with an EMU configuration file that isn't in your current working directory (or named `.emu.yaml`), your command sentence would look like below.
```bash
emu get systems --config /path/to/my/emu-config-file.yaml
```
