CREATE TABLE IF NOT EXISTS Users(
    user_id SERIAL PRIMARY KEY,
    address VARCHAR(256) NOT NULL,
    email  VARCHAR(256) NOT NULL,
    phone VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS Orders(
    order_id SERIAL PRIMARY KEY,
    price INT NOT NULL,
    unit_nums VARCHAR(60000),
    amount VARCHAR(60000),
    state VARCHAR(32),
    user_id INT,

    CONSTRAINT FK_order_user FOREIGN KEY(user_id)
        REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Reviews(
    user_id INT,
    text VARCHAR(60000),

    CONSTRAINT FK_review_user FOREIGN KEY(user_id)
        REFERENCES Users(user_id)
);
