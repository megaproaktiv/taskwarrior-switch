# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-10-29

### Fixed

- Header detection when parsing `todo ls` output is now case-insensitive.
  Headers such as `id`, `Id`, or `ID` are all recognized and skipped, so a
  non-uppercase header row is no longer mistaken for a task row.

### Added

- Case-insensitive header parsing tests covering lowercase and mixed-case headers.

[0.1.0]: https://github.com/megaproaktiv/taskwarrior-switch/releases/tag/v0.1.0
