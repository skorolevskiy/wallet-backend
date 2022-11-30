CREATE TABLE users (
    "id" serial PRIMARY KEY,
    "name" varchar(255) not null,
    "username" varchar(255) not null,
    "password" varchar(255) not null
);

CREATE TABLE wallet (
    "wallet_id" serial PRIMARY KEY,
    "name" varchar(255) not null,
    "user_id" int not null,
    "balance" float not null,
    "currency" varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE transaction (
    "id" serial PRIMARY KEY,
    "wallet_id" int not null,
    "amount" float not null,
    "balance_after" float not null,
    "commission_amount" float,
    "currency" varchar(255) not null,
    "created_at" TIMESTAMP not null,
    FOREIGN KEY (wallet_id) REFERENCES wallet(wallet_id) ON DELETE CASCADE
);
