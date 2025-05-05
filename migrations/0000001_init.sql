CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    gender     VARCHAR(255) NOT NULL,
    birthday   DATE         NOT NULL,
    address    VARCHAR(255),
    phone      VARCHAR(255),
    avatar     VARCHAR(255),
    role       VARCHAR(255) NOT NULL
);

CREATE TABLE books
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    bookname    VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    category    VARCHAR(55) NOT NULL,
    author      VARCHAR(255) NOT NULL,
    stock       INT           NOT NULL,
    price       DECIMAL(10, 2) NOT NULL,
    image       VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE orders
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT NOT NULL,
    status     VARCHAR(255) NOT NULL,
    total      DECIMAL(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE order_details
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    order_id   INT NOT NULL,
    book_id    INT NOT NULL,
    quantity   INT NOT NULL,
    price      DECIMAL(10) NOT NULL,
    total      DECIMAL(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (book_id) REFERENCES books (id)
);

CREATE TABLE payments
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    order_id   INT NOT NULL,
    payer      VARCHAR(255) NOT NULL,
    amount     DECIMAL(10) NOT NULL,
    method     VARCHAR(255) NOT NULL,
    status     VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id)
);