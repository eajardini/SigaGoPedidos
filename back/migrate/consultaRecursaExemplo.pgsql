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




-- Teste para verificar funcionamento
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






-- delete from menu;