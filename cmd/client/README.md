# Triage Client

A simple CLI tool that interacts with the Triage server. A path to a config file is expected under the environment variable `TRIAGE_CLIENT_CONFIG_PATH`, the [sample config](https://github.com/nylar/triage/blob/master/etcd/client.sample.toml) should be a good starting point.

## Tickets

### List

Retrieves a list of all tickets

**Args**:

None

### Create

Creates a new ticket

**Args**:

Arg | Description | Required?
--- | --- | ---
`--subject` | The subject of the ticket | Yes

## Comments

### List

Retrieves a list of all comments for a ticket

**Args**:

Arg | Description | Required?
--- | --- | ---
`--ticket_id` | Find all comments by this ticket ID | Yes

### Create

Creates a new comment for a ticket

**Args**:

Arg | Description | Required?
--- | --- | ---
`--ticket_id` | Ticket ID to associate the comment with | Yes
`--content` | Body of the comment | Yes
