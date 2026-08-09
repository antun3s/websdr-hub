# Station Catalog

## Purpose

Modelo de dados, validação e carregamento de estações WebSDR a partir de arquivos YAML no diretório `data/stations/`. Define o schema de cada estação e garante integridade do catálogo contra duplicatas, erros de digitação e dados inválidos.

## Requirements

### Requirement: Station data model
The system SHALL define a Station struct with fields: id, name, url, software, location (country, city, coordinates), languages, coverage (array of bands), max_users, operator, and added_at.

#### Scenario: Complete station is valid
- **WHEN** a station YAML file contains all required fields with valid values
- **THEN** validation returns zero errors

#### Scenario: Station ID is a kebab-case slug
- **WHEN** a station ID contains uppercase letters, spaces, or special characters other than hyphens
- **THEN** validation SHALL reject it with an error message describing the slug format

#### Scenario: Software is a known value
- **WHEN** a station software field is not one of websdr, kiwisdr, openwebrx, phantomsdr, or other
- **THEN** validation SHALL reject it with an error listing valid values

#### Scenario: Country is ISO 3166-1 alpha-2
- **WHEN** a station country field is not exactly 2 uppercase letters
- **THEN** validation SHALL reject it with an error message

#### Scenario: Coordinates are in valid range
- **WHEN** a station latitude is outside [-90, 90] or longitude is outside [-180, 180], or coordinates are exactly [0, 0]
- **THEN** validation SHALL reject it with an error message indicating the invalid range

#### Scenario: Coverage bands are valid
- **WHEN** a station has no coverage bands, a band has no name, start_hz is negative, or stop_hz is not greater than start_hz
- **THEN** validation SHALL reject it with an error message for each invalid band

#### Scenario: added_at is valid date format
- **WHEN** a station added_at field is not in YYYY-MM-DD format
- **THEN** validation SHALL reject it with an error message

#### Scenario: Validation accumulates all errors
- **WHEN** a station file has multiple validation errors
- **THEN** validation SHALL return all errors at once, not just the first error encountered

### Requirement: YAML loading from directory
The system SHALL load all `.yaml` and `.yml` files from a specified directory and parse each as a Station struct.

#### Scenario: Load all station files from directory
- **WHEN** the loader reads a directory containing valid YAML station files
- **THEN** it SHALL return a sorted list of Station structs (alphabetically by ID) and zero errors

#### Scenario: Filename must match station ID
- **WHEN** a station file is named `nl-enschede.yaml` but contains `id: de-berlin` in its content
- **THEN** validation SHALL reject it with an error indicating filename and ID do not match

#### Scenario: Unknown YAML fields are detected
- **WHEN** a station file contains a YAML key not defined in the Station struct (e.g., a typo like `cordinates`)
- **THEN** the loader SHALL detect it and return an error naming the unknown field

#### Scenario: Duplicate station IDs are detected
- **WHEN** two station files contain the same ID value
- **THEN** the loader SHALL return an error indicating the duplicate ID

#### Scenario: Duplicate URLs are detected
- **WHEN** two station files contain normalized URLs that resolve to the same host and path
- **THEN** the loader SHALL return an error indicating the duplicate URL

### Requirement: URL normalization for duplicate detection
The system SHALL normalize URLs for comparison by lowercasing the host and stripping the trailing slash from the path.

#### Scenario: URLs with trailing slash difference
- **WHEN** two stations have URLs `http://example.com:8073` and `http://example.com:8073/`
- **THEN** the system SHALL treat them as the same URL and report a duplicate

#### Scenario: URLs with different host casing
- **WHEN** two stations have URLs with hosts `EXAMPLE.COM` and `example.com`
- **THEN** the system SHALL treat them as the same URL and report a duplicate