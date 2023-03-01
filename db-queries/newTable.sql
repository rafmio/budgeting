--Create new data base:
CREATE DATABASE budgeting;

-- Switch to budgeting database:
\c budgeting

CREATE TABLE transactions (
    transaction_id serial PRIMARY KEY,
    account VARCHAR (20),
    trdate DATE, -- date of transaction
    trtype VARCHAR (64), -- type of transaction (tecnical or to counterpary)
    docdate DATE, -- date of document
    docnumb VARCHAR (10), -- number of document
    counterpary VARCHAR (64),
    cntp_tax_id VARCHAR (12),
    cntp_contract VARCHAR (64),
    purpose VARCHAR (255), -- purpose of payment
    comnt VARCHAR (255), -- comment himself
    direction VARCHAR (10), --inflow, outflow
    amount NUMERIC (2),
    item VARCHAR (64)
);
