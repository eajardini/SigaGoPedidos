

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