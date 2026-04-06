# `emu`
A tool for automating eMASS records management.

## Quickstart
**Step 1.** Compile the source code and install `emu`. 
```bash
make
```

**Step 2.** Set your eMASS API keys as environment variables. 
```bash
export $EMASS_API_KEY_PRODUCTION="aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
export $EMASS_API_KEY_PILOT="bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
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
emu get artifacts --system-id 1984
```

