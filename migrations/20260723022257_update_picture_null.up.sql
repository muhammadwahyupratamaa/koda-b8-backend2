UPDATE users
SET picture = ''
WHERE picture IS NULL;

ALTER TABLE users
ALTER COLUMN picture SET DEFAULT '';

SELECT id, name, picture
FROM users;