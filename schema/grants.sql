-- see https://github.com/andyxning/shortme#grant
grant select, insert, update, delete on sequence.* to 'sequence'@'%' identified by 'sequence';
grant insert, update on shortme.* to 'shortme_w'@'%' identified by 'shortme_w';
grant select on shortme.* to 'shortme_r'@'%' identified by 'shortme_r';

