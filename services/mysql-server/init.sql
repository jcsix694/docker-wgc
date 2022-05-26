CREATE SCHEMA `wgc-database`;

CREATE USER 'wgc'@'%' IDENTIFIED BY 'root';
GRANT ALL ON *.* to 'wgc'@'%';

SET GLOBAL log_bin_trust_function_creators=1;

flush privileges;