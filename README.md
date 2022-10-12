## Enigma Mart

### Database
```sql
CREATE database enigmart_api;

CREATE TABLE m_product (
    id varchar(60) primary key,
    name varchar(100) not null,
    price int,
    stock int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE t_product_sell (
    id varchar(60) primary key,
    date date,
    factur varchar(60),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE t_product_sell_detail (
    id varchar(60) primary key,
    product_sell_id varchar(60),
    product_id varchar(60),
    qty int,
    price int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    foreign key(product_sell_id) references t_product_sell(id),
    foreign key(product_id) references m_product(id)
);
```

### Run
```
SET DB_HOST=localhost
SET DB_NAME=enigmart_api
SET DB_PORT=5432
SET DB_USER=postgres
SET DB_PASSWORD=YourDBPassword
SET API_URL=:YourAPIPort
go run .
```
