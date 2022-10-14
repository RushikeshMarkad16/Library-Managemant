CREATE TABLE Transaction(
    id VARCHAR(64) NOT NULL,
    issue_date VARCHAR(100) NOT NULL,
    due_date VARCHAR(100) NOT NULL,
    return_date VARCHAR(100),
    book_id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
    ON DELETE CASCADE,
    FOREIGN KEY(book_id) REFERENCES book(id)
    ON DELETE CASCADE
    );