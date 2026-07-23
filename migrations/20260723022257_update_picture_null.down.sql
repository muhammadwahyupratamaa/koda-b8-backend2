UPDATE users
SET picture = NULL
WHERE picture = '';

ALTER TABLE users
ALTER COLUMN picture DROP DEFAULT;

