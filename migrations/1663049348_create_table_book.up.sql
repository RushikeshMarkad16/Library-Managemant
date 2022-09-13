CREATE TABLE book
(
    id VARCHAR(64) NOT NULL,
    name VARCHAR(20) NOT NULL,
    author VARCHAR(20) NOT NULL,
    price INT NOT NULL,
    total_copies INT NOT NULL,
    status VARCHAR(10) NOT NULL,
    available_copies INT NOT NULL,
    PRIMARY KEY (id)
);
