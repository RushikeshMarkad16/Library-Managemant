CREATE TABLE Transaction(
    id VARCHAR(64) NOT NULL,
    issue_date VARCHAR(40) NOT NULL,
    return_date VARCHAR(40) NOT NULL,
    actual_return_date VARCHAR(40) NOT NULL,
    book_id VARCHAR(40) NOT NULL,
    user_id VARCHAR(40) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(book_id) REFERENCES book(id)
    );