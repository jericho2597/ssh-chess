#!/bin/sh

# Check if the database file exists
if [ ! -f /data/db/chess.db ]; then
    echo "Initializing database..."
    sqlite3 /data/db/chess.db <<EOF
CREATE TABLE IF NOT EXISTS games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    white TEXT,
    black TEXT,
    result TEXT,
    moves TEXT,
    date DATETIME DEFAULT CURRENT_TIMESTAMP
);
EOF
    echo "Database initialized successfully."
else
    echo "Database already exists. Skipping initialization."
fi

# Keep the container running
tail -f /dev/null