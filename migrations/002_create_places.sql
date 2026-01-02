-- Migration: Create places table
-- Created: Dia 10

CREATE TABLE IF NOT EXISTS places (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    city_id VARCHAR(255) NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    lng DOUBLE PRECISION NOT NULL,
    profiles TEXT[] NOT NULL,
    description TEXT,
    makes_sense_for TEXT,
    does_not_make_sense_if TEXT,
    tags TEXT[],
    FOREIGN KEY (city_id) REFERENCES cities(id) ON DELETE CASCADE
);

CREATE INDEX idx_places_city_id ON places(city_id);
CREATE INDEX idx_places_slug ON places(slug);

