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