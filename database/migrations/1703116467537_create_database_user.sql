CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';

-- Important: The password used in this migration is for demonstration purposes only.
-- You should use a different, randomly generated password for your own applications.
-- Due to the fact that this code is PUBLIC AVAILABLE and the password is VISIBLE
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';