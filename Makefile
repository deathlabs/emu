# ---------------------------------------------------------
# Set default values.
# ---------------------------------------------------------

.DEFAULT_GOAL := build
VERSION := $(shell date +%Y-%m-%d-%H%M%S)
SBOM_FILE ?= sbom.json
SBOM_SCAN_FILE ?= scan.json
SBOM_SCAN_FAILURE_THRESHOLD ?= low
VEX_YAML_FILE ?= vex.yaml
VEX_FILE ?= vex.json
VEX_AUTHOR ?= Victor Fernandez III
VEX_ID ?= emu

# ---------------------------------------------------------
# Generate the VEX document.
# ---------------------------------------------------------

define VEX_FILTER
{
  "@context": "https://openvex.dev/ns/v0.2.0",
  "@id": strenv(VEX_ID),
  "author": strenv(VEX_AUTHOR),
  "timestamp": strenv(VEX_TIMESTAMP),
  "version": 1,
  "statements": [
    .advisories[] | {
      "vulnerability": {
        "name": .vulnerability
      },
      "products": [
        .products[] | {
          "@id": .
        }
      ],
      "status": .status,
      "justification": .justification,
      "impact_statement": .impact_statement
    }
  ]
}
endef

export VEX_FILTER

.PHONY: vex
.SILENT: vex
vex:
	echo "[*] Generating $(VEX_FILE) from $(VEX_YAML_FILE)"
	VEX_ID="$(VEX_ID)" \
	VEX_AUTHOR="$(VEX_AUTHOR)" \
	VEX_TIMESTAMP="$$(date -u +'%Y-%m-%dT%H:%M:%SZ')" \
	yq --output-format=json "$$VEX_FILTER" "$(VEX_YAML_FILE)" > "$(VEX_FILE)"

# ---------------------------------------------------------
# Generate the SBOM.
# ---------------------------------------------------------

.PHONY: sbom
.SILENT: sbom
sbom:
	echo "[*] Generating $(SBOM_FILE)"
	syft dir:. --output "cyclonedx-json=$(SBOM_FILE)"

# ---------------------------------------------------------
# Scan the SBOM.
# ---------------------------------------------------------

.PHONY: scan
.SILENT: scan
scan: sbom vex
	echo "[*] Scanning $(SBOM_FILE)"
	grype "sbom:$(SBOM_FILE)" \
		--vex "$(VEX_FILE)" \
		--fail-on "$(SBOM_SCAN_FAILURE_THRESHOLD)" \
		--output table \
		--output "cyclonedx-json=$(SBOM_SCAN_FILE)"

# ---------------------------------------------------------
# Build the artifact.
# ---------------------------------------------------------

.PHONY: build
.SILENT: build

build: scan
	go install -ldflags="-s -w -X 'github.com/deathlabs/emu/v4/cmd.version=$(VERSION)'" . 

# ---------------------------------------------------------
# Update dependencies.
# ---------------------------------------------------------

.PHONY: update
.SILENT: update

update:
	go get -u ./...
	go mod tidy
