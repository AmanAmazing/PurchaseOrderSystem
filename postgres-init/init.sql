--- Users table -----
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(100),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE departments (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE user_departments (
    user_id INTEGER REFERENCES users(id),
    department_id INTEGER REFERENCES departments(id),
    PRIMARY KEY (user_id,department_id)
);

CREATE TABLE roles(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE user_roles (
    user_id INTEGER REFERENCES users(id),
    role_id INTEGER REFERENCES roles(id),
    PRIMARY KEY (user_id,role_id)
);

CREATE TYPE priority_level AS ENUM ('low','medium','high');
CREATE TYPE purchase_order_approval AS ENUM ('rejected','pending','manager_approved','partial_approved','approved');
CREATE TABLE purchase_orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    department_id INTEGER REFERENCES departments(id),
    title VARCHAR(255),
    description TEXT,
    status purchase_order_approval DEFAULT 'pending',
    priority priority_level DEFAULT 'low',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TYPE approval_status AS ENUM ('rejected','pending','improve','manager_approved','approved');
CREATE TABLE purchase_order_items (
    id SERIAL PRIMARY KEY,
    purchase_order_id INTEGER REFERENCES purchase_orders(id),
    item_name VARCHAR(255) NOT NULL,
    supplier VARCHAR(255) NOT NULL,
    nominal VARCHAR(255) NOT NULL,
    product VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    link VARCHAR(255) NOT NULL,
    status approval_status DEFAULT 'pending',
    approver_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE purchase_order_approvers (
    id SERIAL PRIMARY KEY, 
    user_id INTEGER REFERENCES users(id)
);

-------------  Currently working on --------------------
CREATE TABLE suppliers ( 
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE nominals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
