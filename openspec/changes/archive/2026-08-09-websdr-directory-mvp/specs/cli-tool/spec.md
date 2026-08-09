## ADDED Requirements

### Requirement: Validate subcommand
The system SHALL provide a `validate` subcommand that reads all station files from `data/stations/`, runs validation, and reports all errors to stdout.

#### Scenario: Validate reports all errors
- **WHEN** `validate` runs against a directory with multiple invalid station files
- **THEN** all validation errors SHALL be printed to stdout and the command SHALL exit with a non-zero status code

#### Scenario: Validate succeeds on clean catalog
- **WHEN** `validate` runs against a directory where all station files are valid
- **THEN** no errors SHALL be printed and the command SHALL exit with code 0

### Requirement: Build subcommand
The system SHALL provide a `build` subcommand that loads the station catalog and generates `dist/v1/stations.json`.

#### Scenario: Build generates stations.json
- **WHEN** `build` runs with a valid catalog
- **THEN** `dist/v1/stations.json` SHALL be created and the command SHALL exit with code 0

### Requirement: Check subcommand
The system SHALL provide a `check` subcommand that runs health checks against all stations in the catalog, merges results with previous status, and generates `dist/v1/status.json`.

#### Scenario: Check with previous status
- **WHEN** `check` runs with `-prev dist/v1/status.json` and 3 stations in the catalog
- **THEN** each station SHALL be probed via HTTP, results merged with previous status, and `dist/v1/status.json` SHALL be written atomically

#### Scenario: Check with no previous status
- **WHEN** `check` runs without a valid previous status source (file not found or HTTP unavailable)
- **THEN** the system SHALL start from empty state and still produce a valid `dist/v1/status.json`

#### Scenario: Check marks offline stations without failing
- **WHEN** `check` runs and some stations are offline
- **THEN** the command SHALL exit with code 0 (offline is data, not a failure)

### Requirement: Check subcommand flags
The system SHALL support the following flags on the `check` subcommand: `-concurrency` (int, default 10), `-timeout` (duration per attempt, default 10s), `-retry` (duration between attempts, default 5s), `-vantage` (string, observation point ID), and `-prev` (string, path or URL to previous status).

#### Scenario: Custom concurrency
- **WHEN** `check` runs with `-concurrency 5`
- **THEN** a maximum of 5 HTTP requests SHALL be in flight simultaneously

#### Scenario: Custom timeout
- **WHEN** `check` runs with `-timeout 15s`
- **THEN** each HTTP probe SHALL have a 15-second timeout

#### Scenario: Vantage point identification
- **WHEN** `check` runs with `-vantage gh-actions-ewr1`
- **THEN** the generated status JSON SHALL include `"vantage": "gh-actions-ewr1"`

### Requirement: CLI builds to a single binary
The system SHALL compile to a single static binary named `websdrctl` from the `cmd/websdrctl` package.

#### Scenario: Binary is buildable
- **WHEN** `go build -o bin/websdrctl ./cmd/websdrctl` is executed
- **THEN** a runnable binary SHALL be produced at `bin/websdrctl`