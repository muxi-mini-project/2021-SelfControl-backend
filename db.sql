Create database SC;
Use SC

DROP TABLE IF EXISTS users;
create table users(
  student_id       varchar(100)	not null ,
  password         varchar(100) not null ,
  user_picture     varchar(100) not null ,
  gold             int          not null ,
  name             varchar(100) not null ,
  privacy          int          not null ,
  current_backdrop int          not null ,
  primary key (student_id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS gold_histories;
create table gold_histories(
  id              int          not null auto_increment ,
  student_id	    varchar(100) not null ,
  time			      varchar(100) not null ,
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
  time       varchar(100) not null ,
  day        int          not null ,
  month      int          not null ,
  primary key (id)
)ENGINE=InnoDB; 

DROP TABLE IF EXISTS punch_contents;
create table punch_contents(
  id          int          not null ,
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

DROP TABLE IF EXISTS users_backdrops;
create table users_backdrops(
  id          int          not null auto_increment ,
  student_id  varchar(100) not null ,
  backdrop_id int          not null ,
  primary key (id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS week_lists;
create table week_lists(
  id         int          not null auto_increment ,
  ranking    int          not null ,
  student_id varchar(100) not null ,
  number     int          not null ,
  day        int          not null ,
  primary key (id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS month_lists;
create table month_lists(
  id         int          not null auto_increment ,
  ranking    int          not null ,
  student_id varchar(100) not null ,
  number     int          not null ,
  month      int          not null ,
  primary key (id)
)ENGINE=InnoDB;

DROP TABLE IF EXISTS list_histories;
create table list_histories(
  id         int          not null auto_increment ,
  type       int          not null ,
  student_id varchar(100) not null ,
  former     int          not null ,
  after      int          not null ,
  primary key (id)
)ENGINE=InnoDB;