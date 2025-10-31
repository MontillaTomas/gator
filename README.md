# Gator 🐊

An RSS feed aggregator CLI tool built in Go. Gator (aggreGATOR) allows you to collect RSS feeds from across the internet, store posts in a PostgreSQL database, and browse aggregated content directly from your terminal.

## Features

- Add and manage RSS feeds from any source
- Store collected posts in PostgreSQL
- Follow and unfollow feeds added by any user
- Browse recent posts with links to full content
- Multi-user support with authentication

## Prerequisites

Before installing Gator, ensure you have:

- **Go** (version 1.16 or higher) - [Download Go](https://golang.org/dl/)
- **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install Gator using `go install`:

```bash
go install github.com/yourusername/gator@latest
```

Make sure your `$GOPATH/bin` is in your system's PATH to run the `gator` command.

## Configuration

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://<username>:<password>@<host>:<port>/<database_name>?sslmode=disable",
  "current_user_name": ""
}
```

Replace the placeholders with your PostgreSQL connection details:
- `<username>`: Your PostgreSQL username (e.g., `postgres`)
- `<password>`: Your PostgreSQL password
- `<host>`: Database host (e.g., `localhost`)
- `<port>`: PostgreSQL port (default: `5432`)
- `<database_name>`: Name of your database (e.g., `gator`)

Example:
```json
{
  "db_url": "postgres://postgres:mypassword@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Usage

### Getting Started

1. **Register a new user**
   ```bash
   gator register <username>
   ```
   Creates a new user and sets them as the current user in the config file.

2. **Login as an existing user**
   ```bash
   gator login <username>
   ```
   Switches to an existing user account.

### Managing Feeds

3. **Add a new feed**
   ```bash
   gator addfeed <feed_name> <feed_url>
   ```
   Adds a new RSS feed to the database and automatically follows it. Requires login.

4. **View all feeds**
   ```bash
   gator feeds
   ```
   Lists all feeds in the database.

5. **Follow a feed**
   ```bash
   gator follow <feed_url>
   ```
   Follow an existing feed. Requires login.

6. **Unfollow a feed**
   ```bash
   gator unfollow <feed_url>
   ```
   Unfollow a feed you're currently following. Requires login.

7. **View your followed feeds**
   ```bash
   gator following
   ```
   Shows all feeds you're currently following. Requires login.

### Aggregating and Browsing Posts

8. **Aggregate posts**
   ```bash
   gator agg
   ```
   Fetches and stores all available posts from feeds in the database.

9. **Browse recent posts**
   ```bash
   gator browse [limit]
   ```
   Displays the most recent posts from feeds you follow. Default limit is 2 posts if not specified. Requires login.
   
   Example:
   ```bash
   gator browse 10
   ```

### User Management

10. **List all users**
    ```bash
    gator users
    ```
    Shows all registered users in the database.

11. **Reset database**
    ```bash
    gator reset
    ```
    ⚠️ **Warning**: Deletes all users from the database. Use with caution.

## Example Workflow

```bash
# Register and login
gator register alice

# Add some feeds
gator addfeed "Go Blog" https://go.dev/blog/feed.atom
gator addfeed "Tech News" https://example.com/rss

# Aggregate posts
gator agg

# Browse latest posts
gator browse 5

# Manage follows
gator following
gator unfollow https://example.com/rss
```