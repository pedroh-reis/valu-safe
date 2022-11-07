INSERT INTO locker (id, unlocked, timestamp)
VALUES
    (0, TRUE, '2022-11-07 07:00:00'),
    (0, FALSE, '2022-11-07 07:01:00'),
    (0, TRUE, '2022-11-07 07:02:00'),
    (0, False, '2022-11-07 07:03:00');


INSERT INTO locker (id, unlocked, timestamp)
VALUES
    (0, TRUE, NOW());