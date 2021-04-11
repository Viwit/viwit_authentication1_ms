CREATE DATABASE viwit;

USE viwit;

CREATE TABLE users (
    user_id INT(12) AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR (30) NOT NULL,
    lastname VARCHAR (30) NOT NULL,
    email VARCHAR (50) NOT NULL,
    reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    user_password VARCHAR (128) NOT NULL,
    wallet_id INT(12),
    token_id INT (12)
);
CREATE TABLE logins(
    login_id INT (32) AUTO_INCREMENT PRIMARY KEY,
    user_id INT(12) NOT NULL,
    login_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users ( firstname,lastname,email,user_pasword)
VALUES (
    'Jairo','Suarez','jaasuarezvi@unal.edu.co','asd123'
);