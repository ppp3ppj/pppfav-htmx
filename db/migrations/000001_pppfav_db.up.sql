BEGIN;

-- Set timezone
SET TIME ZONE 'asia/bangkok';
-- Install uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

COMMIT;
