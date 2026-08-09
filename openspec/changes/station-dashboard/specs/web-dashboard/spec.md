## ADDED Requirements

### Requirement: Fetch and merge catalog and status data
The system SHALL fetch `stations.json` and `status.json` using relative URLs from the same origin on page load, merge them by station ID, and render the combined data.

#### Scenario: Both endpoints respond successfully
- **WHEN** the page loads and both `./v1/stations.json` and `./v1/status.json` return valid JSON
- **THEN** the system SHALL display a list of stations with catalog info (name, country, software, bands) and status info (online/offline, last checked, latency)

#### Scenario: Status endpoint is unavailable
- **WHEN** the page loads and `stations.json` succeeds but `status.json` fails
- **THEN** the system SHALL display all stations with status marked as "unknown" and a note that status data is temporarily unavailable

#### Scenario: Both endpoints fail
- **WHEN** the page loads and both fetches fail
- **THEN** the system SHALL display an error message with a retry button

#### Scenario: Loading state
- **WHEN** the page is loading data from the API
- **THEN** the system SHALL display a loading indicator until both fetches complete

### Requirement: Filter by online status
The system SHALL allow filtering stations by online status: all, online only, or offline only.

#### Scenario: Filter online only
- **WHEN** the user selects "Online" in the status filter
- **THEN** only stations with `state: "online"` in the status data SHALL be displayed

#### Scenario: Filter offline only
- **WHEN** the user selects "Offline" in the status filter
- **THEN** only stations with `state: "offline"` or unknown status SHALL be displayed

#### Scenario: Show all
- **WHEN** the user selects "All" in the status filter
- **THEN** all stations SHALL be displayed regardless of status

### Requirement: Filter by software type
The system SHALL allow filtering stations by software type using checkboxes.

#### Scenario: Select one software type
- **WHEN** the user checks "WebSDR" in the software filter
- **THEN** only stations using WebSDR software SHALL be displayed

#### Scenario: Select multiple software types
- **WHEN** the user checks both "WebSDR" and "OpenWebRX"
- **THEN** stations using either WebSDR or OpenWebRX SHALL be displayed

#### Scenario: No software selected
- **WHEN** the user unchecks all software types
- **THEN** all stations SHALL be displayed regardless of software

### Requirement: Filter by country
The system SHALL allow filtering stations by country via a dropdown populated from available countries in the catalog.

#### Scenario: Select a country
- **WHEN** the user selects "NL" from the country dropdown
- **THEN** only stations located in the Netherlands SHALL be displayed

#### Scenario: Clear country filter
- **WHEN** the user selects "All countries" from the country dropdown
- **THEN** all stations SHALL be displayed regardless of country

### Requirement: Filter by frequency band
The system SHALL allow filtering stations by frequency band via a dropdown populated from available bands in the catalog.

#### Scenario: Filter by HF
- **WHEN** the user selects "HF" from the band dropdown
- **THEN** only stations that cover the HF band SHALL be displayed

#### Scenario: Filter by VHF
- **WHEN** the user selects "2m" from the band dropdown
- **THEN** only stations that cover the 2m band SHALL be displayed

### Requirement: Filter by language
The system SHALL allow filtering stations by interface language via a dropdown populated from available languages in the catalog.

#### Scenario: Filter by language
- **WHEN** the user selects "pt" from the language dropdown
- **THEN** only stations that support Portuguese SHALL be displayed

### Requirement: Filters are combinable
The system SHALL apply all active filters simultaneously (AND logic).

#### Scenario: Combine status and country filters
- **WHEN** the user selects status "Online" and country "BR"
- **THEN** only stations that are BOTH online AND in Brazil SHALL be displayed

#### Scenario: Combine software and band filters
- **WHEN** the user checks "KiwiSDR" in software and selects "HF" in band
- **THEN** only KiwiSDR stations that cover HF SHALL be displayed

### Requirement: Active filter count indicator
The system SHALL display the number of active filters and the number of stations matching the current filters.

#### Scenario: Filters active
- **WHEN** the user has applied 2 filters and 5 stations match
- **THEN** the system SHALL display "5 estações (2 filtros ativos)" or equivalent message

#### Scenario: No filters
- **WHEN** no filters are active
- **THEN** the system SHALL display only the total station count without filter indication

### Requirement: Station card display
Each station SHALL be displayed as a card showing: name, software badge, country flag emoji, city, frequency bands with modes, operator callsign, and online/offline status indicator.

#### Scenario: Online station card
- **WHEN** a station is online
- **THEN** the card SHALL show a green status indicator, latency in ms, and "checked at" timestamp

#### Scenario: Offline station card
- **WHEN** a station is offline
- **THEN** the card SHALL show a red status indicator, consecutive failure count, and last online timestamp

### Requirement: Responsive layout
The system SHALL use CSS grid to display station cards in a responsive layout: 1 column on narrow screens, 2 columns on medium screens, 3 or more columns on wide screens.

#### Scenario: Mobile viewport
- **WHEN** the viewport width is less than 640px
- **THEN** station cards SHALL be displayed in a single column

#### Scenario: Desktop viewport
- **WHEN** the viewport width is greater than 1024px
- **THEN** station cards SHALL be displayed in 3 or more columns

### Requirement: No external dependencies
The system SHALL be a single self-contained HTML file with embedded CSS and JavaScript. It SHALL NOT require any external libraries, frameworks, CDN resources, or build tools.

#### Scenario: Open in browser directly
- **WHEN** a user opens `index.html` in a modern browser with an internet connection to the same origin
- **THEN** the dashboard SHALL render correctly without any npm install or build step

### Requirement: Network origin independence
The system SHALL use only relative URLs (e.g., `./v1/stations.json`) for all API requests, never hardcoding an absolute origin, hostname, or port. The dashboard SHALL function identically regardless of how the hosting server is reached — via `localhost`, `127.0.0.1`, LAN IP address, public domain name, custom port, or GitHub Pages URL.

#### Scenario: Accessed via localhost
- **WHEN** the dashboard is served from `http://localhost:8080/`
- **THEN** the API fetches SHALL use relative paths resolving to `http://localhost:8080/v1/stations.json` without any origin mismatch

#### Scenario: Accessed via LAN IP
- **WHEN** the dashboard is served from `http://192.168.1.50:3000/`
- **THEN** the API fetches SHALL use relative paths resolving to `http://192.168.1.50:3000/v1/status.json` without any origin mismatch

#### Scenario: Accessed via public domain
- **WHEN** the dashboard is served from `https://example.github.io/websdr-directory/`
- **THEN** the API fetches SHALL use relative paths resolving to `https://example.github.io/websdr-directory/v1/stations.json` without any origin mismatch