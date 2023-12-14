CREATE TABLE customers
(
    uuid VARCHAR(256) PRIMARY KEY,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    email VARCHAR(256) NOT NULL
);

CREATE TABLE subscription_plans
(
    uuid VARCHAR(256) PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

CREATE TABLE subscriptions
(
    uuid VARCHAR(256) PRIMARY KEY,
    expire_date DATETIME NOT NULL,
    customer VARCHAR(256) NOT NULL,
    subscription_plan VARCHAR(256) NOT NULL,
    FOREIGN KEY(customer) REFERENCES customers(uuid)
    FOREIGN KEY(subscription_plan) REFERENCES subscription_plans(uuid)
);
