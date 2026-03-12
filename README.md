# Seer CLI - Search & Discovery Examples

This document provides examples for using the search and discovery commands in the Seer CLI.

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
