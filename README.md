# Gator CLI

A CLI app for aggregating feeds, built with Go and PostgreSQL. This tool allows users to register, follow feeds and browse posts.

## Installation

To get started, you need to have [Go](https://go.dev/doc/install) and [PostgreSQL](https://www.postgresql.org/download/) installed.

1. Clone the repository

```bash
git clone https://github.com/syeero7/blog-aggregator.git gator
cd gator
```

2. Create .gatorconfig.json file in the home directory

```bash
echo '{ "db_url": <postgres_database_url> }' > ~/.gatorconfig.json
```

3. Run migrations and generate queries

```bash
# install goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# run migrations
cd sql/schema
goose postgres <database_url> up
cd ../..

# install sqlc 
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# generate queries
sqlc generate 
```

4. Install gator

```bash
go install
```

## Usage

| Command | Description | Example |
| ------------------------ | -------------------------- | ------------------ |
| `gator register [username]` | Registers a new user | `gator register jack` |
| `gator login [username]` | Login an existing user | `gator login jack` |
| `gator add-feed [feed_name] [feed_url]` | Adds a new feed | `gator "Boot.dev Blog" https://blog.boot.dev/index.xml` |
| `gator feeds` | Lists all feeds | `gator feeds` |
| `gator follow [feed_url]` | Follows a feed by its URL | `gator follow https://blog.boot.dev/index.xml` |
| `gator unfollow [feed_url]` | Unfollows a feed by its URL | `gator unfollow https://blog.boot.dev/index.xml` |
| `gator users` | Lists all registered users | `gator users` |
| `gator following` | Lists all feeds that current user is following | `gator following` |
| `gator agg [time_between_requests]` | Aggregates posts from all followed feeds. `time_between_requests` is a duration string (eg: `10s`,`5m`,`1h`) | `gator agg 30m` |
| `gator browse [limit]` | Lists published posts. `limit` is optional and defaults to 2 | `gator browse 20` |
