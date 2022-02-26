```sql

CREATE TABLE books (
    isbn char(14) NOT NULL,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    category varchar(255) NOT NULL,
    price decimal(5,2) NOT NULL
);
```

```sql
INSERT INTO books (title, author, category, price) VALUES
( 'Mrityunjaya', 'Shivaji Sawant', 'Novel', 944.50),
('Panipat', 'Vishwas Patil', 'Novel',  599),
( 'The Alchemist', 'Paulo Coelho', 'Fiction',699),
('Sapiens: A Brief History of Humankind','Yuval Noah Harari','Non-Fiction',400);

```

```sql
ALTER TABLE books ADD PRIMARY KEY (isbn);
```