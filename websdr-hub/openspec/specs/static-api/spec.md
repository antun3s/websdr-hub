# Static API

## Purpose

Geração dos arquivos JSON estáticos — catálogo de estações (`stations.json`) e status de disponibilidade (`status.json`) — na estrutura `dist/v1/` pronta para publicação via GitHub Pages.

## Requirements

### Requirement: Generate stations catalog JSON
The system SHALL serialize the loaded station catalog to a JSON file at `dist/v1/stations.json`.

#### Scenario: Stations JSON is generated
- **WHEN** the build command is executed with a valid station catalog
- **THEN** a `stations.json` file SHALL be created at `dist/v1/stations.json` containing all stations serialized as JSON

#### Scenario: Empty catalog produces empty array
- **WHEN** the catalog contains zero stations
- **THEN** the output JSON SHALL contain `[]` (empty array)

### Requirement: Output directory structure
The system SHALL place all generated static API files under `dist/v1/` to allow API versioning.

#### Scenario: Output directory is created
- **WHEN** the build or check command runs for the first time and `dist/v1/` does not exist
- **THEN** the system SHALL create the directory before writing output files

### Requirement: JSON output fields match station model
The system SHALL include all Station struct fields in the generated JSON, using the same field names and structure as the YAML input model.

#### Scenario: Station JSON includes all fields
- **WHEN** a station has coverage bands with modes and a location with coordinates
- **THEN** the output JSON SHALL include `id`, `name`, `url`, `software`, `location` (with `country`, `city`, `coordinates`), `languages`, `coverage` (with `name`, `start_hz`, `stop_hz`, `modes`), `max_users`, `operator`, and `added_at`