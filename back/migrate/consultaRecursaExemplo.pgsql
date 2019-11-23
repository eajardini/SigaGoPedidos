--drop table menu
CREATE TABLE menu (
   menuID               int,
   codigo               varchar(15),
   nivel                int,
   descricao            varchar(40),
   icone                varchar(30),
   codigoMenuSuperiorID int
);

INSERT INTO menu (menuID,codigo, nivel,descricao, codigoMenuSuperiorID, icone
)
VALUES
   (1,  'vendas',1, 'Vendas', NULL, NULL),
   (2,  'vendas',2, 'Clientes', 1, NULL),
   (3,  'vendas',3, 'Cadastro', 2, NULL),
   (4,  'vendas',3, 'Relatóriso', 2, NULL),
   (5,  'financeiro',1, 'Financeiro', null, NULL),
   (6,  'financeiro',2, 'Contas a Pagar', 5, NULL),
   (7,  'financeiro',3, 'Manutenção', 6, NULL),
   (8,  'financeiro',3, 'Relatórios', 6, NULL),
   (9,  'financeiro',2, 'Contas a Receber', 5, NULL),
   (10, 'financeiro',3, 'Relatórios', 9, NULL);

WITH RECURSIVE submenus AS (
   SELECT  menuID, descricao, nivel, codigoMenuSuperiorID, CAST(descricao As varchar(1000)) As Itens_Menu
   FROM    menu  
   where codigoMenuSuperiorID is null
   UNION
   SELECT  m.menuID, m.descricao, m.nivel, m.codigoMenuSuperiorID, CAST(s.Itens_Menu || '->' || m.descricao As varchar(1000)) As Itens_Menu
   FROM menu m
   INNER JOIN submenus s ON s.menuID = m.codigoMenuSuperiorID
) 
SELECT  descricao, nivel, Itens_Menu
  FROM submenus
ORDER BY Itens_Menu;

WITH RECURSIVE supplytree AS (
     SELECT si_id, si_item, si_parentid, CAST(si_item As varchar(1000)) As si_item_fullname
     FROM supplyitem
     WHERE si_parentid IS NULL
     UNION ALL
     SELECT si.si_id,si.si_item, si.si_parentid, CAST(sp.si_item_fullname || '->' || si.si_item As varchar(1000)) As si_item_fullname
     FROM supplyitem As si
	   INNER JOIN supplytree AS sp
	         ON (si.si_parentid = sp.si_id)
)
SELECT si_id, si_item_fullname
FROM supplytree
ORDER BY si_item_fullname;

-- delete from menu;