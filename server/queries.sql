CREATE TABLE Users (
id serial PRIMARY KEY,
mobile_num VARCHAR (15) UNIQUE NOT NULL,
profile_pic VARCHAR (100),
created_on TIMESTAMP NOT NULL
);

CREATE TABLE Otp (
id serial PRIMARY KEY,
mobile_num VARCHAR (15) UNIQUE NOT NULL,
otp VARCHAR (6),
created_on TIMESTAMP NOT NULL,
last_otp    TIMESTAMP NOT NULL
);

CREATE TABLE groups (
	id serial PRIMARY KEY,
	users integer[],
	admin_id integer,
	created_on TIMESTAMP NOT NULL
)

INSERT INTO otp (mobile_num, otp, created_on, last_otp) VALUES('9825447696', '111111', current_timestamp, current_timestamp) 
				ON CONFLICT (mobile_num) 
				DO UPDATE 
				SET
					otp = '111111', 
					last_otp = current_timestamp

