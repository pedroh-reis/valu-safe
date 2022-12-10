INSERT INTO locker (id, unlocked, timestamp)
VALUES
    (0, TRUE, NOW() - INTERVAL '12 hour'),
    (0, FALSE, NOW() - INTERVAL '10 hour'),
    (0, TRUE, NOW() - INTERVAL '5 hour'),
    (0, FALSE, NOW() - INTERVAL '2 hour'),

    (1, TRUE, NOW() - INTERVAL '12 hour'),
    (1, FALSE, NOW() - INTERVAL '10 hour'),
    (1, TRUE, NOW() - INTERVAL '5 hour'),
    (1, FALSE, NOW() - INTERVAL '2 hour');