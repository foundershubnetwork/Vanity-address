-- Create vanity_keypairs table
CREATE TABLE IF NOT EXISTS vanity_keypairs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    public_key VARCHAR(255) UNIQUE NOT NULL,
    secret_key_base58 TEXT NOT NULL,
    secret_key BYTEA NOT NULL,
    isUsed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    used_at TIMESTAMP,
    vanity_pattern VARCHAR(50),
    generation_attempts INTEGER
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_vanity_keypairs_isUsed ON vanity_keypairs(isUsed);
CREATE INDEX IF NOT EXISTS idx_vanity_keypairs_public_key ON vanity_keypairs(public_key);