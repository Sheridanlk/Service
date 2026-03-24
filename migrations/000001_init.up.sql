CREATE TABLE IF NOT EXISTS user_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'active',
    connection_data JSONB,
    instance_id INTEGER REFERENCES instances(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS instances (
    id SERIAL PRIMARY KEY,
    protocol VARCHAR(20) NOT NULL,
    port INTEGER NOT NULL CHECK (port > 0 AND port <= 65535),
    local_ip INET NOT NULL,
    settings JSONB,
    is_active BOOLEAN DEFAULT true,
    UNIQUE(local_ip, port)
);