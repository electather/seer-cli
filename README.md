# seer-cli

A command-line interface for the [Seer](https://github.com/seerr/app) media request management API.

## Installation

### Linux / macOS

```sh
curl -fsSL https://raw.githubusercontent.com/electather/seer-cli/main/install.sh | sh
```

Installs the latest stable release to `/usr/local/bin` (uses `sudo` if needed, falls back to `~/.local/bin`). Supports `amd64` and `arm64`.

### Manual

Download the archive for your platform from the [Releases](https://github.com/electather/seer-cli/releases) page, extract it, and move the binary to your `PATH`.

---

## Configuration

Set your server URL and API key once:

```sh
seer-cli config set --server https://your-seer-instance.com --api-key YOUR_KEY
```

Or pass them as flags on every command:

```sh
seer-cli --server https://your-seer-instance.com --api-key YOUR_KEY <command>
```

---

## Global Flags
All commands support the following global flags:
- `-s, --server`: Seer server URL (e.g., `http://localhost:5055`).
- `-k, --api-key`: Your Seer API key.
- `-v, --verbose`: Enable verbose output (shows request URL and HTTP status).

---

## 1. Multi Search
Search for movies, TV shows, and people in a single query.

**Example: Search for "The Matrix"**
```bash
seer-cli search multi -q "The Matrix"
```

**Example: Search for "Christopher Nolan" with page 2**
```bash
seer-cli search multi -q "Christopher Nolan" --page 2
```

---

## 2. Keyword Search
Search for TMDB keywords.

**Example: Search for "sci-fi"**
```bash
seer-cli search keyword -q "sci-fi"
```

---

## 3. Company Search
Search for production companies.

**Example: Search for "Warner Bros."**
```bash
seer-cli search company -q "Warner Bros."
```

---

## 4. Trending
Get trending movies and TV shows.

**Example: Get current trending items**
```bash
seer-cli search trending
```

**Example: Get trending items for the week (if supported by server)**
```bash
seer-cli search trending --time-window week
```

---

## 5. Discover Movies
Discover and filter movies using various criteria.

**Example: Discover Drama movies (Genre ID 18)**
```bash
seer-cli search movies --genre 18
```

**Example: Discover movies from a specific studio (Studio ID 1)**
```bash
seer-cli search movies --studio 1
```

**Example: Discover movies sorted by release date**
```bash
seer-cli search movies --sort-by primary_release_date.desc
```

---

## 6. Discover TV Shows
Discover and filter TV shows using various criteria.

**Example: Discover Drama TV shows (Genre ID 18)**
```bash
seer-cli search tv --genre 18
```

**Example: Discover TV shows from a specific network (Network ID 1)**
```bash
seer-cli search tv --network 1
```
