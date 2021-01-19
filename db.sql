
DROP TABLE IF EXISTS users;
create table users(
  student_ID int		  not null ,
  gold       int          not null ,
  name       varchar(100) not null ,
  primary key (student_ID)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS gold_histories;
create table gold_histories(
  student_ID	  int          not null ,
  time			  varchar(100) not null ,
  change_number   int          not null ,
  residual_number int		   not null ,
  reason          varchar(200) not null ,
  primary key (student_ID)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS backdrops;
create table backdrops(
  backdrop_ID int          not null auto_increment ,
  picture_URL varchar(100) not null ,
  price       int          not null ,
  primary key (backdrop_ID)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS punch_histories;
create table punch_histories(
  student_ID int          not null ,
  punch_ID   int          not null ,
  time       varchar(100) not null ,
  primary key (student_ID,punch_ID,time)
)ENGINE=InnoDB; 

DROP TABLE IF EXISTS achievements;
create table achievements(
  student_ID  int not null ,
  achievement text ,
  primary key (student_ID)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS punch_contents;
create table punch_contents(
  punch_ID    int          not null auto_increment ,
  content     text ,
  picture_URL varchar(100) not null ,
  primary key (punch_ID)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS users_punchs;
create table users_punchs(
  student_ID int not null ,
  punch_ID   int not null ,
  primary key (student_ID,punch_ID)
)ENGINE=InnoDB;



