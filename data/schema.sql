CREATE DOMAIN eth_addr AS bytea CHECK(octet_length(value) = 20);
CREATE DOMAIN sha3_256 AS bytea CHECK(octet_length(value) = 32);

CREATE TABLE services (
    id sha3_256 PRIMARY KEY,
    so json NOT NULL
);

CREATE TABLE vpn_services (
    service_id sha3_256 PRIMARY KEY REFERENCES services (id)
    out_speed_kibs int NOT NULL,
    in_speed_kibs int NOT NULL
);

CREATE TABLE clients (
    id eth_addr PRIMARY KEY,
    added timestamp NOT NULL
);

CREATE TABLE payments (
    id sha3_256 PRIMARY KEY,
    service_id sha3_256 NOT NULL REFERENCES services (id),
    client_id eth_addr NOT NULL REFERENCES clients (id),
    eth_block int NOT NULL,
    solt bigint NOT NULL,
    password sha3_256 NOT NULL
);

CREATE TABLE vpn_payments (
    payment_id sha3_256 PRIMARY KEY REFERENCES payments (id),
    total_mibs int NOT NULL
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY,
    payment_id sha3_256 NOT NULL REFERENCES payments (id),
    started timestamp NOT NULL,
    ended timestamp NOT NULL
);

CREATE TABLE vpn_sessions (
    session_id uuid PRIMARY KEY REFERENCES sessions (id),
    consumed_mibs int NOT NULL
);
