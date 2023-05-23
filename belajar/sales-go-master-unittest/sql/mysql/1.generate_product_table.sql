create table product (
	id int not null auto_increment,
	name varchar(100) not null,
	price float not null,
	created_at datetime,
	updated_at datetime,
	deleted_at datetime,
	primary key (id)
)