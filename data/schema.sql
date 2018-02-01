CREATE DOMAIN eth_addr AS bytea CHECK(octet_length(value) = 20);
CREATE DOMAIN keccak_256 AS bytea CHECK(octet_length(value) = 32);

CREATE TABLE services (
    id keccak_256 PRIMARY KEY,
    so json NOT NULL
);

CREATE TABLE vpn_services (
    service_id keccak_256 PRIMARY KEY REFERENCES services (id)
);

CREATE TABLE clients (
    id eth_addr PRIMARY KEY
);

CREATE TABLE payments (
    id keccak_256 PRIMARY KEY,
    service_id keccak_256 NOT NULL REFERENCES services (id),
    client_id eth_addr NOT NULL REFERENCES clients (id),
    eth_block int NOT NULL,
    solt bigint NOT NULL,
    password keccak_256 NOT NULL
);

CREATE TABLE vpn_payments (
    payment_id keccak_256 PRIMARY KEY REFERENCES payments (id),
    total_mibs int NOT NULL
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY,
    payment_id keccak_256 NOT NULL REFERENCES payments (id),
    started timestamp NOT NULL,
    ended timestamp
);

CREATE TABLE vpn_sessions (
    session_id uuid PRIMARY KEY REFERENCES sessions (id),
    consumed_mibs int
);
