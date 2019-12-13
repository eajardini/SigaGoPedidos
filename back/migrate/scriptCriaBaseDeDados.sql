CREATE USER usigagopedidos WITH PASSWORD 'p0stdb@';

create database sigagopedidosdev ENCODING "utf-8";
ALTER DATABASE sigagopedidosdev SET datestyle TO ISO, DMY;

\c sigagopedidosdev

GRANT CONNECT ON DATABASE sigagopedidosdev TO usigagopedidos;
GRANT USAGE ON SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO usigagopedidos;
grant ALL on database sigagopedidosdev to usigagopedidos;

create database sigagopedidosteste ENCODING "utf-8";
ALTER DATABASE sigagopedidosteste SET datestyle TO ISO, DMY;

\c sigagopedidosteste

GRANT CONNECT ON DATABASE sigagopedidosteste TO usigagopedidos;
GRANT USAGE ON SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO usigagopedidos;
grant ALL on database sigagopedidosteste to usigagopedidos;

create database sigagopedidosprod  ENCODING "utf-8";
ALTER DATABASE sigagopedidosprod SET datestyle TO ISO, DMY;


\c sigagopedidosprod

GRANT CONNECT ON DATABASE sigagopedidosprod TO usigagopedidos;
GRANT USAGE ON SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usigagopedidos;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO usigagopedidos;
grant ALL on database sigagopedidosprod to usigagopedidos;




