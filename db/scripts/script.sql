CREATE DATABASE viwit;
USE viwit;

CREATE TABLE users (
    user_id INT(12) AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR (30) NOT NULL,
    lastname VARCHAR (30) NOT NULL,
    email VARCHAR (50) NOT NULL,
    reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    user_pasword VARCHAR (128) NOT NULL,
    wallet_id INT(12),
    block_account BOOL NOT NULL,
    user_type INT(2)
);

CREATE TABLE tokens (
    token_id INT (12) AUTO_INCREMENT PRIMARY KEY,
    token VARCHAR (128) NOT NULL,
    user_id INT (12),
    expiration_date TIMESTAMP NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    foreign key (user_id) references users (user_id)
);

CREATE TABLE logins (
    login_id INT (32) AUTO_INCREMENT PRIMARY KEY,
    user_id INT(12) NOT NULL,
    login_date TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

