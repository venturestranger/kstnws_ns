create table notifications (
	id BIGSERIAL PRIMARY KEY,
	id_user BIGINT NOT NULL,
	status INT NOT NULL,
	content VARCHAR(1000) NOT NULL,
	date TIMESTAMP NOT NULL
);
