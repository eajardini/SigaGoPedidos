-- drop table if exists menu
CREATE TABLE if not exists menu (
   menuID               int,
   codigo               varchar(15),
   nivel                int,
   descricao            varchar(40),
   icone                varchar(30),
   codigoMenuSuperiorID int,
   constraint pk_menu primary key(menuID)
);

INSERT INTO menu (menuID,codigo, nivel,descricao, codigoMenuSuperiorID, icone
)
VALUES
   (1,  'vnd100',1, 'Vendas', NULL, NULL),
      (2,  'vnd110',2, 'Clientes', 1, NULL),
          (3,  'vnd111',3, 'Cadastro', 2, NULL),
          (4,  'vnds112',3, 'Relatórios', 2, NULL),
      (11, 'vnd100',2, 'Pedidos', 1, NULL),
          (12, 'vnd121',3, 'Realizar Pedidos', 11, NULL),
          (13, 'vnd122',3, 'Relatórios de Pedidos', 11, NULL),
  (5,  'fncr100',1, 'Financeiro', null, NULL),
      (6,  'fncr110',2, 'Contas a Pagar', 5, NULL),
          (7,  'fncr111',3, 'Manutenção', 6, NULL),
          (8,  'fncr112',3, 'Relatórios', 6, NULL),
      (9,  'fncr120',2, 'Contas a Receber', 5, NULL),
          (10, 'fncr121',3, 'Relatórios', 9, NULL)
   ON CONFLICT DO NOTHING;