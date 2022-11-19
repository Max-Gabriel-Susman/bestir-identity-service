-- migrate up 
CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    balance INT8
);

-- migrate down 
-- DROP TABLE accounts;