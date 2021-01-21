
DROP TABLE IF EXISTS users;
create table users(
  student_id   int		    not null UNIQUE,
  password     varchar(100) not null ,
  user_picture varchar(100) ,
  gold         int          not null ,
  name         varchar(100) ,
  primary key (student_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS gold_histories;
create table gold_histories(
  student_id	  int          not null ,
  time			  date         not null ,
  change_number   int          not null ,
  residual_number int		   not null ,
  reason          varchar(200) not null ,
  primary key (student_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS backdrops;
create table backdrops(
  backdrop_id int          not null auto_increment ,
  picture_url varchar(100) not null UNIQUE ,
  price       int          not null ,
  primary key (backdrop_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS punch_histories;
create table punch_histories(
  id         int          not null auto_increment ,
  student_id int          not null ,
  title      varchar(100) not null ,
  time       date         not null ,
  primary key (id)
)ENGINE=InnoDB; 

DROP TABLE IF EXISTS achievements;
create table achievements(
  student_id  int not null ,
  achievement text ,
  primary key (student_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS punch_contents;
create table punch_contents(
  --punch_id    int          not null auto_increment ,
  type        varchar(100) not null ,
  title       varchar(100) not null UNIQUE ,
  content     text ,
  picture_url varchar(100) not null UNIQUE ,
  primary key (title)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS users_punchs;
create table users_punchs(
  id         int          not null auto_increment ,
  student_id int          not null ,
  title      varchar(100) not null ,
  primary key (id)
)ENGINE=InnoDB;



