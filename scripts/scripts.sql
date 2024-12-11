-- Create the database
CREATE DATABASE gowtdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

-- Create user if it doesn't exist and grant privileges
CREATE USER 'samuelarthur'@'localhost' IDENTIFIED BY 'itssurabaya';
GRANT ALL PRIVILEGES ON gowtdb.* TO 'samuelarthur'@'localhost';
FLUSH PRIVILEGES;

-- Switch to the database
USE gowtdb;

-- Create the categories table
CREATE TABLE categories (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(80) COLLATE utf8_unicode_ci NOT NULL,
    description text COLLATE utf8_unicode_ci,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Create the tools table
CREATE TABLE tools (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
    category_id int(11) NOT NULL,
    url varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    rating int(11) DEFAULT NULL,
    notes text COLLATE utf8_unicode_ci,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;