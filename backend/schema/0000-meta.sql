CREATE DATABASE tinylibrary;
CREATE USER tinylibrary WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE tinylibrary to tinylibrary;
DROP TABLE books;
DROP TABLE patrons;
DROP TABLE loans;
