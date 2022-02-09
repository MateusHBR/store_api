
CREATE TABLE IF NOT EXISTS products(
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  removed_at TIMESTAMP
);