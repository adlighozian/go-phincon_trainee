create table product (
	id serial not null,
	name varchar(100) not null,
	price float not null,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	primary key (id)
)