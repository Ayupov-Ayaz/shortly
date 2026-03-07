CREATE TABLE IF NOT EXISTS urls (
    id BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(10) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE, 

    CONSTRAINT short_code_length CHECK (char_length(short_code) BETWEEN 3 AND 10)
);

CREATE INDEX idx_urls_short_code ON urls(short_code);
CREATE INDEX idx_urls_created_at ON urls(created_at);
CREATE INDEX idx_urls_active ON urls(active) WHERE active = true;