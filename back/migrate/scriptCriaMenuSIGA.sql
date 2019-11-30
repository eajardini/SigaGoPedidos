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
   (1,  'vendas',1, 'Vendas', NULL, NULL),
      (2,  'vendas',2, 'Clientes', 1, NULL),
          (3,  'vendas',3, 'Cadastro', 2, NULL),
          (4,  'vendas',3, 'Relatórios', 2, NULL),
      (11, 'vendas',2, 'Pedidos', 1, NULL),
          (12, 'vendas',3, 'Realizar Pedidos', 11, NULL),
          (13, 'vendas',3, 'Relatórios de Pedidos', 11, NULL),
  (5,  'financeiro',1, 'Financeiro', null, NULL),
      (6,  'financeiro',2, 'Contas a Pagar', 5, NULL),
          (7,  'financeiro',3, 'Manutenção', 6, NULL),
          (8,  'financeiro',3, 'Relatórios', 6, NULL),
      (9,  'financeiro',2, 'Contas a Receber', 5, NULL),
          (10, 'financeiro',3, 'Relatórios', 9, NULL)
   ON CONFLICT DO NOTHING;