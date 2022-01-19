SELECT s.id,s.user_name,u.user_name as parent_user_name 
FROM stock.user s
left JOIN stock.user u ON u.id = s.parent order by id asc;