CREATE TABLE IF NOT EXISTS contacts(
    id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL,
    position VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS tasks(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    status VARCHAR(16) NOT NULL,
    priority VARCHAR(32) NOT NULL,
    created_by VARCHAR(64),
    created_at TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL
);