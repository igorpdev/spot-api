-- Migration: Create cities table
-- Created: Dia 10

CREATE TABLE IF NOT EXISTS cities (
    id VARCHAR(255) PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE INDEX idx_cities_slug ON cities(slug);

