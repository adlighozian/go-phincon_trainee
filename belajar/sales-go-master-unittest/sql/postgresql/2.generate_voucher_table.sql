create table voucher (
	id serial not null,
	code varchar(100) not null,
	persen float not null,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	primary key (id)
)