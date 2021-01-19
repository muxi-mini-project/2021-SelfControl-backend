create table users(
  student_ID int		  not null ,
  gold       int          not null ,
  name       vatchar(100) not null ,
  primary key (student_ID)
);

create table gold_histories(
  student_ID	  int          not null ,
  time:			  vatchar(100) not null ,
  change_number   int          not null ,
  residual_number int		   not null ,
  reason          vatchar(200) not null ,
  primary key (student_ID)
);

create table backdrops(
  backdrop_ID int          not null auto_increment ,
  picture_URL vatchar(100) not null ,
  price       int          not null ,
  primary key (backdrop_ID)
);

create table punch_histories(
  student_ID int          not null ,
  punch_ID   int          not null ,
  time       vatchar(100) not null ,
  primary key (student_ID,punch_ID,time)
); 

create table achievements(
  student_ID  int          not null ,
  achievement vatchar(200) not null ,
  primary key (student_ID)
);

create table punch_contents(
  punch_ID    int          not null auto_increment ,
  content     vatchar(200) not null ,
  picture_URL vatchar(100) not null ,
  primary key (punch_ID)
);

create table users_punchs(
  student_ID int not null ,
  punch_ID   int not null ,
  primary key (student_ID,punch_ID)
);



