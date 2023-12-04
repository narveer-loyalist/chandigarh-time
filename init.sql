-- Create the database if not exists
CREATE DATABASE IF NOT EXISTS torontotime;

-- Use the database
USE torontotime;

-- Create the time_table table if not exists
CREATE TABLE IF NOT EXISTS time_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
