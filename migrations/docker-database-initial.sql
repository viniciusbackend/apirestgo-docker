create table produtos (
    id serial primary key,
    nome varchar,
    tipo varchar,
    valor Real,
    quantidade int
);

INSERT INTO produtos(nome, tipo, valor, quantidade) VALUES
('fanta uva', 'refrigerante', 7, 10) 