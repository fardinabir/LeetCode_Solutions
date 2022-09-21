select a.id Id from Weather a
join Weather b
on DATEDIFF(a.recordDate, b.recordDate)=1 and a.Temperature>b.Temperature;