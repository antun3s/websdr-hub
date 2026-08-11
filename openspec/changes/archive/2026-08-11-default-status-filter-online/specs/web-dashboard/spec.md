## MODIFIED Requirements

### Requirement: Filter by online status
The system SHALL allow filtering stations by online status: all, online only, or offline only. The status filter SHALL default to "online" on page load.

#### Scenario: Filter online only
- **WHEN** the user selects "Online" in the status filter
- **THEN** only stations with `state: "online"` in the status data SHALL be displayed

#### Scenario: Filter offline only
- **WHEN** the user selects "Offline" in the status filter
- **THEN** only stations with `state: "offline"` or unknown status SHALL be displayed

#### Scenario: Show all
- **WHEN** the user selects "All" in the status filter
- **THEN** all stations SHALL be displayed regardless of status

#### Scenario: Default to online on page load
- **WHEN** the page loads
- **THEN** the status filter SHALL show "Online" selected by default and only online stations SHALL be displayed