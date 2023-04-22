create table transaction_detail (
	id serial not null,
	transaction_id int not null,
	item varchar(100) not null,
	price float not null,
	quantity int not null,
	total float not null,
	primary key (id),
	foreign key (transaction_id) references transaction(id)
)