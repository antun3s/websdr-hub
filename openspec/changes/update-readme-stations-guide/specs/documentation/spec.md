## ADDED Requirements

### Requirement: README focuses on station contribution onboarding
The project README SHALL serve as the primary entry point for contributors who want to add a new station to the catalog.

#### Scenario: First-time visitor reads the README
- **WHEN** a first-time visitor opens the repository README
- **THEN** they SHALL immediately understand what the project is
- **THEN** they SHALL see a link to the live demo at https://websdr.antunes.pro/
- **THEN** they SHALL find clear instructions on how to add a new station

### Requirement: Tutorial for adding a station
The README SHALL contain a step-by-step tutorial for adding a new station to the catalog.

#### Scenario: Contributor adds a station
- **WHEN** a contributor follows the tutorial in the README
- **THEN** they SHALL learn to create a `<id>.yaml` file in `data/stations/`
- **THEN** they SHALL learn to run `./bin/websdrctl validate`
- **THEN** they SHALL learn to open a Pull Request

### Requirement: Data model reference in README
The README SHALL include a concise data model reference showing the YAML structure for a station entry.

#### Scenario: Contributor checks field format
- **WHEN** a contributor reads the data model section
- **THEN** they SHALL see each field name, type, and example value
- **THEN** the example SHALL use a real-world station as illustration

### Requirement: Infrastructure details moved to CONTRIBUTING.md
Details about health check scheduling, GitHub Actions pipeline, license information, and CLI flags SHALL be documented in CONTRIBUTING.md, not in README.md.

#### Scenario: Developer needs infrastructure details
- **WHEN** a developer needs details about health checks or CI/CD
- **THEN** they SHALL find that information in CONTRIBUTING.md
- **THEN** the README SHALL reference CONTRIBUTING.md for advanced topics