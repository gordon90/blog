 create table article(
    -> id int(11) unsigned auto_increment,
    -> title varchar(255),
    -> content text,
    -> keyword varchar(255),
    -> create_time int(20),
    -> primary key (id)
    -> )engine=Innodb default charset=utf8;
