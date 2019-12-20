--drop table if exists menu

--delete from menu where menuID >= 19;

CREATE TABLE if not exists menu (
   menuID               int,
   codigo               varchar(15),
   nivel                int,
   descricao            varchar(40),
   icone                varchar(30),
   link                 varchar(40),               
   codigoMenuSuperiorID int,
   constraint pk_menu primary key(menuID)
);

INSERT INTO menu (menuID,codigo, nivel,descricao, codigoMenuSuperiorID, link, icone 
)
VALUES
   (1,  'vnd100',1, 'Vendas', 0, 'sem link', ''),     
   (2,  'vnd110',2, 'Clientes', 1, 'sem link', ''),
          (3,  'vnd111',3, 'Cadastro', 2, '/cadastroclientes',''),
          (4,  'vnds112',3, 'Relatórios', 2, '/relatorioclientes',''),
      (11, 'vnd100',2, 'Pedidos', 1, 'sem link', ''),
          (12, 'vnd121',3, 'Realizar Pedidos', 11,'/realizarpedidos' ,''),
          (13, 'vnd122',3, 'Relatórios de Pedidos', 11,'/relatoriopedidos' ,''),
  (5,  'fncr100',1, 'Financeiro', 0, 'sem link', ''),
      (6,  'fncr110',2, 'Contas a Pagar', 5, 'sem link', ''),
          (7,  'fncr111',3, 'Manutenção', 6, '/cadastrocontaspagar' ,''),
          (8,  'fncr112',3, 'Relatórios', 6, '/relatoriocontaspagar' ,''),
      (9,  'fncr120',2, 'Contas a Receber', 5, 'sem link', ''),
          (10, 'fncr121',3, 'Relatórios', 9, '/relatoriocontasreceber' ,''),
          (14, 'fncr122',3, 'Mensagens', 9, '/messages' ,''),
  (15,  'stqu100',1, 'Estoque', 0, 'sem link', ''),
      (16,  'fncr110',2, 'Produto', 15, 'sem link', ''),
          (17,  'fncr111',3, 'Manutenção', 16, '/cadastroproduto' ,''),
          (18,  'fncr112',3, 'Relatórios', 16, '/relatorioproduto' ,''),
  (19,  'admn100',1, 'Administração', 0, 'sem link', ''),
      (20,  'admn110',2, 'Usuários', 19, 'sem link', ''),
          (21,  'admn111',3, 'Manutenção', 20, '/cadastrousuario' ,''),
          (22,  'admn112',3, 'Relatórios', 20, '/relatoriousuario' ,'') ,
      (23,  'admn120',2, 'Grupos', 19, 'sem link', ''),
          (24,  'admn111',3, 'Manutenção', 23, '/cadastrogrupo' ,''),
          (25,  'admn112',3, 'Relatórios', 23, '/relatoriogrupo' ,'') 
   
   ON CONFLICT DO NOTHING;

   -- Consulta que retorna o Menu do sistema
WITH RECURSIVE submenus AS (
   SELECT  menuID, descricao, nivel, link,icone,codigoMenuSuperiorID, CAST(descricao As varchar(1000)) As Itens_Menu
   FROM    menu  
   where codigoMenuSuperiorID = 0
   UNION
   SELECT  m.menuID, m.descricao, m.nivel, m.link, m.icone,m.codigoMenuSuperiorID, CAST(s.Itens_Menu || '->' || m.descricao As varchar(1000)) As Itens_Menu
   FROM menu m
   INNER JOIN submenus s ON s.menuID = m.codigoMenuSuperiorID
) 
SELECT  menuID, codigoMenuSuperiorID, descricao, link , icone  ,nivel
  FROM submenus
ORDER BY Itens_Menu;

   commit;

