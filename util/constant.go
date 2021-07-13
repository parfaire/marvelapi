package util

import "time"

// redis
const REDIS_KEY_ALL_CHARACTER_IDS = "marvel-all-character-ids"
const REDIS_KEY_MARVEL_TTL = 24 * time.Hour

// others
const MARVEL_API_PORT = 8080
const MARVEL_MAX_GET = 100

// error message
const ERROR_NOT_FOUND = "Not Found."
const ERROR_INTERNAL_SERVER = "Internal Server Error."
