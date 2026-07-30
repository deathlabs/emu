# Changelog

## [4.0.4] - 2026-07-28
### Added
- Added `emu get cloud-resources`
- Added `emu get cloud-resource-results`

## [4.0.3] - 2026-07-27
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

## [4.0.2] - 2026-05-24
### Added
- Added `emu upload artifact`

### Changed
- Converted the `contentType` parameter into a map-based `headers` parameter in the `emass.POST` function

### Removed
- Removed comments that weren't adding value

## [4.0.1] - 2026-05-24
### Added
- Added `emu test api` subcommand for testing connectivity to the eMASS API

### Changed
- Replaced `Vic Fernandez III` with `Victor Fernandez III` throughout

## [4.0.0] - 2026-04-23
### Added
- Added initial v4 release

### Changed
- Migrated from v3 configuration format

[4.0.4]: https://github.com/deathlabs/emu/releases/tag/v4.0.4
[4.0.3]: https://github.com/deathlabs/emu/releases/tag/v4.0.3
[4.0.2]: https://github.com/deathlabs/emu/releases/tag/v4.0.2
[4.0.1]: https://github.com/deathlabs/emu/releases/tag/v4.0.1
[4.0.0]: https://github.com/deathlabs/emu/releases/tag/v4.0.0
