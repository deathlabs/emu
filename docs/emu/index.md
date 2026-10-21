# EMU
Enterprise Mission Assurance Support Service (eMASS) Updater (EMU) is a tool for automating eMASS records management.

## Quickstart

**Step 1.** First, compile the source code into an executable program using the provided `Makefile`.
```bash
make
```

**Step 2.** Then, set the environment variables `emu` needs to run.
```
export EMASS_API_URL="https://connect.emass.apps.mil"
export EMASS_API_UID="1234567890"
export EMASS_API_KEY="aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
export EMASS_SYSTEM_ID="1234"
export NPE_CERTIFICATE=$(cat SomeApp.cer)
export NPE_KEY=$(cat SomeApp.pem)
```
