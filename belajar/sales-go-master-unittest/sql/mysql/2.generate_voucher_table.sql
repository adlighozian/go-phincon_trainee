create table voucher (
	id int not null auto_increment,
	code varchar(100) not null,
	persen float not null,
	created_at datetime,
	updated_at datetime,
	deleted_at datetime,
	primary key (id)
)