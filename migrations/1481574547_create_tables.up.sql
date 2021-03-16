CREATE TABLE employees
(
    id              INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at      DATE,
    updated_at      DATE,
    deleted_at      DATE,
    name            VARCHAR(30),
    address         VARCHAR(255),
    dob             DATE,
    salary_in_cents INT,
    department      VARCHAR(50),
    date_of_joining DATE
);
