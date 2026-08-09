## MODIFIED Requirements

### Requirement: Filter by frequency
The system SHALL provide an input field where the user can enter a frequency in kHz. The system SHALL filter stations to only show those whose coverage bands include the entered frequency (where `start_hz ≤ freq × 1000 ≤ stop_hz` for at least one band).

#### Scenario: Filter by 40m frequency
- **WHEN** the user enters "7100" in the frequency filter
- **THEN** only stations that have a coverage band covering 7,100,000 Hz SHALL be displayed

#### Scenario: Frequency with no matches
- **WHEN** the user enters "99999999" in the frequency filter
- **THEN** no stations SHALL be displayed and the counter SHALL show "0 stations"

#### Scenario: Empty frequency clears filter
- **WHEN** the user clears the frequency input field
- **THEN** all stations SHALL be displayed (filter removed)