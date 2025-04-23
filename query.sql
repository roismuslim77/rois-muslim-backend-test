select murid.id as id_murid, murid.name, p.status as pendidikan_terakhir, murid.time_create, p.time_create as time_update from murid
join pendidikan p on murid.id = p.id_murid
LEFT OUTER JOIN pendidikan p2 ON (murid.id = p2.id_murid AND
    (p.time_create < p2.time_create OR (p.time_create = p2.time_create AND p.id < p2.id)))
where p2.id is null 