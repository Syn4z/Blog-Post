create table articles (
  id serial not null unique,
  title varchar(64),
  content text,
  primary key(id)
);


insert into articles(title, content)
values
    ('Hello World', 'The obligatory Hello World Article ...'),
    ('Another Article', 'Yet another blog article about something exciting');