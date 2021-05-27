CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL,
    active BOOLEAN NOT NULL 
);

CREATE TABLE users (
    userid serial PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL
)