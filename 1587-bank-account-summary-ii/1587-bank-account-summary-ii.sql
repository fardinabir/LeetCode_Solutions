SELECT U.name, SUM(T.amount) AS 'balance'
FROM Transactions T
INNER JOIN Users U
ON T.account = U.account
GROUP BY U.name
HAVING SUM(T.amount) >= 10000