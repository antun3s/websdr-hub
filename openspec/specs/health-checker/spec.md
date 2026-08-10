# Health Checker

## Purpose

Motor de health check HTTP com retry, controle de concorrência e coleta de latência/status code. Verifica se estações WebSDR estão respondendo, com etiqueta de rede (User-Agent identificado, limite de redirects, timeout curto).

## Requirements

### Requirement: HTTP health probe
The system SHALL perform a single HTTP GET request to a station URL with a configurable timeout, read up to 4KB of the response body, and collect the HTTP status code and latency in milliseconds.

#### Scenario: Station responds with HTTP 200
- **WHEN** a station URL returns HTTP 200 within the timeout
- **THEN** the probe SHALL record online=true, the HTTP status code, and the measured latency

#### Scenario: Station responds with non-200 status
- **WHEN** a station URL returns an HTTP error status (e.g., 502, 503) within the timeout
- **THEN** the probe SHALL record online=false, the HTTP status code, and an error message containing the status code

#### Scenario: Station host is unreachable
- **WHEN** a station URL resolves to a host that does not respond (DNS failure, connection refused, or timeout)
- **THEN** the probe SHALL record online=false and an error message describing the network failure

### Requirement: Retry logic with two attempts
The system SHALL attempt up to 2 HTTP probes per station, with a configurable interval between attempts. If the first attempt succeeds, the second SHALL be skipped. If both attempts fail, the station SHALL be marked offline.

#### Scenario: First attempt succeeds
- **WHEN** the first HTTP probe returns HTTP 200
- **THEN** the system SHALL NOT make a second attempt and SHALL mark the station online with 1 attempt recorded

#### Scenario: First attempt fails, second succeeds
- **WHEN** the first HTTP probe returns 502 and the second returns 200 after the retry interval
- **THEN** the system SHALL mark the station online with 2 attempts recorded and error cleared

#### Scenario: Both attempts fail
- **WHEN** both HTTP probes fail within their respective timeouts
- **THEN** the system SHALL mark the station offline with 2 attempts recorded and an error message from the last attempt

### Requirement: Concurrency control
The system SHALL limit the maximum number of simultaneous HTTP requests using a configurable concurrency cap, defaulting to 10.

#### Scenario: Requests are throttled to concurrency cap
- **WHEN** there are 12 stations to check and concurrency is set to 3
- **THEN** at no point during execution SHALL there be more than 3 HTTP requests in flight simultaneously

#### Scenario: Default concurrency is 10
- **WHEN** no concurrency value is specified
- **THEN** the system SHALL use a maximum of 10 simultaneous HTTP requests

### Requirement: User-Agent identification
The system SHALL include a User-Agent header in every health check request that identifies the project name, version, and repository URL.

#### Scenario: User-Agent header is present
- **WHEN** a health check HTTP request is sent
- **THEN** the request SHALL include a User-Agent header containing `websdr-hub` and the repository URL

### Requirement: HTTP client configuration
The system SHALL use an HTTP client limited to 5 redirects and with a configurable per-request timeout.

#### Scenario: Excessive redirects are stopped
- **WHEN** a station URL redirects more than 5 times (e.g., captive portal loop)
- **THEN** the HTTP client SHALL stop following redirects and return an error