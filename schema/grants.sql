-- see https://github.com/andyxning/shortme#grant
create user if not exists 'sequence'@'%' identified by 'sequence';
create user if not exists 'shortme_w'@'%' identified by 'shortme_w';
create user if not exists 'shortme_r'@'%' identified by 'shortme_r';
grant select, insert, update, delete on sequence.* to 'sequence'@'%';
grant insert, update on shortme.* to 'shortme_w'@'%';
grant select on shortme.* to 'shortme_r'@'%';

