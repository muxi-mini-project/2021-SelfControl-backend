Use Self_Control

DROP TABLE IF EXISTS users;
create table users(
  student_id   varchar(100)	not null UNIQUE,
  password     varchar(100) not null ,
  user_picture varchar(100) not null ,
  gold         int          not null ,
  name         varchar(100) not null ,
  privacy      boolean      not null ,
  primary key (student_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS gold_histories;
create table gold_histories(
  id              int          not null auto_increment ,
  student_id	    varchar(100) not null ,
  time			      datetime     not null ,
  change_number   int          not null ,
  residual_number int		       not null ,
  reason          varchar(200) not null ,
  primary key (id)
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
  student_id varchar(100) not null ,
  title      varchar(100) not null ,
  time       datetime     not null ,
  day        int          not null ,
  month      int          not null ,
  primary key (id)
)ENGINE=InnoDB; 

DROP TABLE IF EXISTS punch_contents;
create table punch_contents(
  id          int          not null auto_increment ,
  type        varchar(100) not null ,
  title       varchar(100) not null UNIQUE ,
  content     varchar(100) ,
  primary key (id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS users_punches;
create table users_punches(
  id         int          not null auto_increment ,
  student_id varchar(100) not null ,
  title      varchar(100) not null ,
  number     int          not null default 0,
  primary key (id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS list_prices;
create table list_prices(
  ranking varchar(100) not null ,
  price   int          not null ,
  primary key (ranking)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS users_backdrops;
create table users_backdrops(
  id          int          not null auto_increment ,
  student_id  varchar(100) not null ,
  backdrop_id varchar(100) not null ,
  primary key (id)
)ENGINE=InnoDB;