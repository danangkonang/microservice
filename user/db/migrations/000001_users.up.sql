CREATE TABLE IF NOT EXISTS users(
	user_id VARCHAR (225) NOT NULL,
	user_name VARCHAR (225) UNIQUE NOT NULL,
	password VARCHAR (225) NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);
