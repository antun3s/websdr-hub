## ADDED Requirements

### Requirement: Load previous status from file or URL
The system SHALL load the previous status file from either a local path or an HTTP URL. If the source is unavailable, it SHALL start with an empty status.

#### Scenario: Load from local file
- **WHEN** the previous status source is a valid local JSON file
- **THEN** the system SHALL parse it into a status File struct

#### Scenario: Load from HTTP URL
- **WHEN** the previous status source is an HTTP URL returning valid JSON
- **THEN** the system SHALL fetch and parse it into a status File struct

#### Scenario: Source is unavailable
- **WHEN** the previous status source (local or HTTP) is unreachable or returns invalid data
- **THEN** the system SHALL start with an empty status (no prior entries) and continue without error

### Requirement: Merge health check results with previous status
The system SHALL merge fresh health check results with the previous status, updating each station's state and preserving state not present in the current run.

#### Scenario: Station is online — reset failures
- **WHEN** a station is found online in the current check
- **THEN** the system SHALL set its consecutive_failures to 0 and update last_online to the current check timestamp

#### Scenario: Station is offline — increment failures
- **WHEN** a station is found offline in the current check and had 3 consecutive_failures in the previous status
- **THEN** the system SHALL set consecutive_failures to 4 and preserve the previous last_online timestamp

#### Scenario: Station is offline — first failure
- **WHEN** a station is found offline in the current check and was not present in the previous status
- **THEN** the system SHALL set consecutive_failures to 1

#### Scenario: Station not checked — preserve previous state
- **WHEN** a station exists in the previous status but was not included in the current check run
- **THEN** the system SHALL preserve its previous state in the output

### Requirement: Status file structure
The system SHALL produce a status file containing a version string, a generated_at timestamp, a vantage point identifier, and a map of station status entries keyed by station ID.

#### Scenario: Status file has required fields
- **WHEN** a status file is generated
- **THEN** it SHALL include `version`, `generated_at`, `vantage`, and `stations` fields

#### Scenario: Each station entry has status fields
- **WHEN** a station entry exists in the status file
- **THEN** it SHALL include `state`, `http_code`, `latency_ms`, `checked_at`, `consecutive_failures`, `last_online`, and `error` fields

### Requirement: Atomic file write
The system SHALL write JSON files to a temporary file in the target directory and rename it atomically to the target path.

#### Scenario: Atomic write prevents partial files
- **WHEN** the system writes a status file
- **THEN** it SHALL first write the complete JSON to a temporary file, then rename to the target path in a single atomic operation