CREATE TABLE IF NOT EXISTS products(
	product_id VARCHAR (225) NOT NULL,
	product_name VARCHAR (225) NOT NULL,
	price INTEGER NOT NULL,
	qty INTEGER NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (product_id)
);
