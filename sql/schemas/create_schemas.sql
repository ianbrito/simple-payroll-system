DROP TABLE IF EXISTS lotations;
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS people;
DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS bonds;

CREATE TABLE IF NOT EXISTS bonds
(
    id         VARCHAR(36)  NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS roles
(
    id         VARCHAR(36)  NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS departments
(
    id         VARCHAR(36)  NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    acronym    VARCHAR(10)  NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS people
(
    id         VARCHAR(36)  NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    cpf        VARCHAR(20)  NOT NULL,
    birth      DATE         NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS employees
(
    id         VARCHAR(36)  NOT NULL PRIMARY KEY,
    enrollment VARCHAR(10)  NOT NULL,
    person_id  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    FOREIGN KEY (person_id) REFERENCES people (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lotations
(
    id             VARCHAR(36)  NOT NULL PRIMARY KEY,
    admission_at   DATE         NOT NULL,
    resignation_at DATE,
    employee_id    VARCHAR(255) NOT NULL,
    department_id  VARCHAR(255) NOT NULL,
    role_id        VARCHAR(255) NOT NULL,
    bond_id        VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    FOREIGN KEY (bond_id) REFERENCES bonds (id) ON DELETE CASCADE
);