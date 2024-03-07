/* 
    This code for creating the function that will be used to update the updated_at timestamp
*/

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

/* 
    This code for creating types table
*/

CREATE TABLE IF NOT EXISTS types (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON types
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


/* 
    This code for creating cars table
*/

CREATE TABLE IF NOT EXISTS cars (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    make VARCHAR(255) NOT NULL,
    model INTEGER NOT NULL,
    color VARCHAR(255) NOT NULL,
    speed INTEGER[2] NOT NULL,
    type_id BIGINT NOT NULL,

    CONSTRAINT fk_types FOREIGN KEY (type_id) REFERENCES types (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON cars
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
