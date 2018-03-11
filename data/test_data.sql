-- Test agent.
INSERT INTO subjects (id, public_key, private_key)
VALUES ('e5b22ad6-16ed-11e8-b642-0ed5f89f718b', '<public_key>',
    '<password_encrypted_private_key>');

-- Test client.
INSERT INTO subjects (id, public_key)
VALUES ('e5b22ed2-16ed-11e8-b642-0ed5f89f718b', '<public_key>');

-- Test service offering.
INSERT INTO offerings (id, agent, service, supply, signature)
VALUES ('e5b231c0-16ed-11e8-b642-0ed5f89f718b',
    'e5b22ad6-16ed-11e8-b642-0ed5f89f718b', 'vpn', 10, '<signature>');

-- Test VPN service offering.
INSERT INTO offerings_vpn (id, hash, country, upload_price, download_price,
    min_upload, max_upload, min_download, max_download, billing_interval,
    billing_deviation, free_intervals, min_upload_bps, min_download_bps,
    template_version, protocol)
VALUES ('e5b231c0-16ed-11e8-b642-0ed5f89f718b',
    '68Ooww1AatJz2wFr60_Scez1SHiaIgPDtrHgn_FKHlk=', 'US',
    'AAAAAAAAAAAAAACrze8BI0VniavN7wAA', 'AAAAAAAAAAAAAACrze8BI0VniavN7wAA',
    1048576, 137438953472, 1048576, 137438953472, 300, 60, 1, 1048576, 2097152,
    1, 'tcp');

-- Test state channel with password: "secret".
INSERT INTO channels (id, agent, client, offering, block, state, total_deposit,
    closed_deposit, salt, password, receipt_balance, receipt_signature)
VALUES (
    'e5b23706-16ed-11e8-b642-0ed5f89f718b',
    'e5b22ad6-16ed-11e8-b642-0ed5f89f718b',
    'e5b22ed2-16ed-11e8-b642-0ed5f89f718b',
    'e5b231c0-16ed-11e8-b642-0ed5f89f718b', 5012227, 'open',
    'AAAAAAAAAACrze-rze8BI0VniavN7wAA', 'AAAAAAAAAAAAAP-rze8BI0VniavN7wAA',
    6012867121110302348, '7U9gC4AZsSZ9E8NabVkw8nHRlFCJe0o_Yh9qMlIaGAg=',
    'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA', '<signature>');
