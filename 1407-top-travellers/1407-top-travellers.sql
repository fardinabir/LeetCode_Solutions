SELECT u.name, ifnull(sum(r.distance),0) as 'travelled_distance' FROM Users u
LEFT join Rides r
ON r.user_id = u.id
GROUP by user_id
ORDER by sum(r.distance) desc, u.name;