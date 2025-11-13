Stocky is a backend service built in Go(Gin) + PostgreSQL

In this service user earns free shares of Indian Stocks as rewards.

Tech Stack
-Go lang
-PostgreSQL
-GORM
-Logrus
-Postman

Database Schema

1.Users
Stores basic user info
id, name, email

2.Rewards
Every time a user gets rewarded shares we add it inside this table
user_id, stock_symbol, shares, received_at, external_ref

3.Holding
Tracks how many total shares a user owns per stock.
user_id, stock_symbol, total_shares

4.Ledger
Tracks the company accounting
-company receives shares from NSE/BSE
-company pays the cash (share value to stock exchange)
-company pays the brokerage fees
entry_group_id, account, debit, credit, shares, stock_symbol
