create table STREAMS (
    ID  serial,
    KEY varchar(64) unique not null,
    VERSION int not null
);