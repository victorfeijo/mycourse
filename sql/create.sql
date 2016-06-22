create table Grade
(
    ID integer not null auto_increment,
    Name varchar (100),
    Completed boolean,
    primary key(ID)
);

create table Note
(
    Name varchar (100),
    Value float,
    GradeID integer,
    foreign key(GradeID) references Grade (ID)
);
