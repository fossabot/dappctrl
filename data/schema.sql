CREATE DOMAIN eth_addr AS char(28); -- RFC-4648 base64 encoded 20 bytes.
CREATE DOMAIN sha3_256 AS char(44); -- RFC-4648 base64 encoded 32 bytes.

CREATE TYPE service_type AS ENUM ('VPN');

CREATE TABLE services (
    id sha3_256 PRIMARY KEY,
    so json NOT NULL,
    type service_type NOT NULL
);

CREATE TABLE vpn_services (
    service_id sha3_256 PRIMARY KEY REFERENCES services (id),
    down_speed_kibs int NOT NULL,
    up_speed_kibs int NOT NULL
);

CREATE TABLE clients (
    id eth_addr PRIMARY KEY,
    added timestamp with time zone NOT NULL
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
    down_mibs int NOT NULL,
    up_mibs int NOT NULL
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY,
    payment_id sha3_256 NOT NULL REFERENCES payments (id),
    server_ip inet,
    client_ip inet,
    client_port int,
    started timestamp with time zone,
    ended timestamp with time zone
);

CREATE TABLE vpn_sessions (
    session_id uuid PRIMARY KEY REFERENCES sessions (id),
    down_kibs int,
    up_kibs int
);
