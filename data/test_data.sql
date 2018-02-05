INSERT INTO services (id, so, type) VALUES (
    '68Ooww1AatJz2wFr60/Scez1SHiaIgPDtrHgn/FKHlk=',
    '{"templateVersion": 1}',
    'VPN'
);

INSERT INTO vpn_services (
    service_id, down_speed_kibs, up_speed_kibs) VALUES (
    '68Ooww1AatJz2wFr60/Scez1SHiaIgPDtrHgn/FKHlk=',
    2048,
    1024
);

INSERT INTO clients (id, added) VALUES (
    'DzmrB7yzddZmi2wCvt0wo1WmMs0=',
    '2018-01-01 00:00:00'
);

INSERT INTO payments (
    id, service_id, client_id, eth_block, solt, password) VALUES (
    'eYElAiJE8fa56RyMyzLepsWvxT4OWfff0ZF7L0xX3M8=',
    '68Ooww1AatJz2wFr60/Scez1SHiaIgPDtrHgn/FKHlk=',
    'DzmrB7yzddZmi2wCvt0wo1WmMs0=',
    5012227,
    6012867121110302348,
    -- Password: secret
    '7U9gC4AZsSZ9E8NabVkw8nHRlFCJe0o/Yh9qMlIaGAg='
);

INSERT INTO vpn_payments (payment_id, down_mibs, up_mibs) VALUES (
    'eYElAiJE8fa56RyMyzLepsWvxT4OWfff0ZF7L0xX3M8=',
    1024,
    256
);
