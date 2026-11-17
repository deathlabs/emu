# Changelog

## [v4.0.7] - 2026-11-17
### Added
- Added user, developer, and CI pipeline documentation with automated GitHub Pages deployment
- Added SBOM generation, OpenVEX generation, and dependency vulnerability scanning to the build workflow

### Changed
- Updated Go modules to remediate dependency vulnerabilities

## [v4.0.6] - 2026-10-1
### Changed
- Updated `golang.org/x/crypto` version
- Updated how the version is set during build time

## [v4.0.5] - 2026-08-19
### Added
- Added `emu get poams`
- Added `emu upload poams`
- Added `emu upload static-code-scan`

## Changed
- The eMASS API guide version used for reference (not all functions were tested against the latest API version).
- Changed `emu upload cloud-resource-results` to `emu upload cloud-resource-scan`

## [v4.0.4] - 2026-08-18
### Added
- Added `emu get cloud-resources`
- Added `emu get cloud-resource-results`
- Added `emu upload cloud-resource-results`

## [v4.0.3] - 2026-07-27
### Added
- Added `emu get applications`
- Added `emu get containers`
- Added `emu get container-scan-results`
- Added `emu get devices`
- Added `emu get device-scan-results`
- Added `emu get static-code-scans`
- Added `emu upload device-scan`

### Removed
- Removed `user-uid` header from HTTP GET requests
- Removed comments that weren't adding value

## [v4.0.2] - 2026-05-24
### Added
- Added `emu upload artifact`

### Changed
- Converted the `contentType` parameter into a map-based `headers` parameter in the `emass.POST` function

### Removed
- Removed comments that weren't adding value

## [v4.0.1] - 2026-05-24
### Added
- Added `emu test api` subcommand for testing connectivity to the eMASS API

### Changed
- Replaced `Vic Fernandez III` with `Victor Fernandez III` throughout

## [v4.0.0] - 2026-04-23
### Added
- Added initial v4 release

### Changed
- Migrated from v3 configuration format

[v4.0.4]: https://github.com/deathlabs/emu/releases/tag/v4.0.4
[v4.0.3]: https://github.com/deathlabs/emu/releases/tag/v4.0.3
[v4.0.2]: https://github.com/deathlabs/emu/releases/tag/v4.0.2
[v4.0.1]: https://github.com/deathlabs/emu/releases/tag/v4.0.1
[v4.0.0]: https://github.com/deathlabs/emu/releases/tag/v4.0.0
