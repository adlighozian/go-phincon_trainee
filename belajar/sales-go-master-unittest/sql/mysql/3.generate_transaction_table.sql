create table transaction (
	id int not null auto_increment,
	transactionnumber int not null unique,
	name varchar(100) not null,
	quantity int not null,
	discount float not null,
	total float not null,
	pay float not null,
	primary key (id)
)