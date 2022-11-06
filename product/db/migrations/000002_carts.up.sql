CREATE TABLE IF NOT EXISTS carts(
	cart_id VARCHAR (225) NOT NULL,
	user_id VARCHAR (225) NOT NULL,
	product_id VARCHAR (225) NOT NULL,
	qty INTEGER NOT NULL,
	is_checkout BOOLEAN,
	CONSTRAINT carts_pkey PRIMARY KEY (cart_id),
	FOREIGN KEY (product_id) REFERENCES products(product_id) ON UPDATE CASCADE
);
