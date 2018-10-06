create table post (
    id int primary key auto_increment,
    thread_id int,
    body text,
    author varchar(80),
    stamp int unique,
    foreign key (thread_id) references thread(id)
);