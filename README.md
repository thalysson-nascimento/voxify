anotações para o banco poder receber os dados devidamente
precisa rodar o comando no psql

`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

para confirmar se esta habilitado

`SELECT * FROM pg_extension WHERE extname = 'uuid-ossp';`
