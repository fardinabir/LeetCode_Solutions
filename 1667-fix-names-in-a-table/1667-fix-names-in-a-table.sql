select user_id, concat(upper(substr(name,1,1)),lower(substr(name,2,length(name)-1))) name from Users order by user_id;