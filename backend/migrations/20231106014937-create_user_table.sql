-- +migrate Up

-- Roles table
CREATE TABLE if NOT EXISTS roles  (
    role_id INT PRIMARY KEY ,
    role_name VARCHAR(50) NOT NULL
);
-- Users table
CREATE TABLE if NOT EXISTS users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

-- Permissions table
CREATE TABLE if NOT EXISTS permissions  (
    permission_id INT PRIMARY KEY ,
    permission_name VARCHAR(50) NOT NULL
);
-- -- UserRoles junction table (many-to-many relationship between users and roles)
-- CREATE TABLE if NOT EXISTS user_roles   (
--     user_id INT,
--     role_id INT,
--     FOREIGN KEY (user_id) REFERENCES users(user_id),
--     FOREIGN KEY (role_id) REFERENCES roles(role_id)
-- );
-- RolePermissions junction table (many-to-many relationship between roles and permissions)
CREATE TABLE if NOT EXISTS role_permissions  (
    role_id INT,
    permission_id INT,
    FOREIGN KEY (role_id) REFERENCES roles(role_id),
    FOREIGN KEY (permission_id) REFERENCES permissions(permission_id)
);
-- UserPermissions table
CREATE TABLE if NOT EXISTS user_permissions  (
    user_id INT,
    permission_id INT,
    is_enabled BOOLEAN DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (permission_id) REFERENCES permissions(permission_id)
);
-- +migrate Down
-- Drop the junction tables first to avoid foreign key constraints
DROP TABLE IF EXISTS user_permissions;
DROP TABLE IF EXISTS role_permissions;
-- DROP TABLE IF EXISTS user_roles;
-- Drop the main tables
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS permissions;

