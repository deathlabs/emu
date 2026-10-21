# Artifacts

## Upload an Artifact
```bash
emu upload artifact \
  -filename policy.pdf
```

## Delete an Artifact
```bash
emu delete artifact \
  -filename policy.pdf
```

## Upload Multiple Files
```bash
emu upload artifacts \
  -filename policies.zip -is-bulk=true
```

## Export an Artifact
```bash
emu get artifact \
  -filename policy.pdf
```
