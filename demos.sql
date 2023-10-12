set odps.namespace.schema=true; 
--查询表dwd_github_events_odps中的100条数据
select * from bigdata_public_dataset.github_events.dwd_github_events_odps where ds = max_pt("bigdata_public_dataset.github_events.dwd_github_events_odps") limit 100;


SET odps.namespace.schema = TRUE;
SELECT dws.repo_id AS repo_id,
 repos.name AS repo_name,
 SUM(dws.stars) AS stars
FROM bigdata_public_dataset.github_events.dws_overview_by_repo_month dws
JOIN bigdata_public_dataset.github_events.db_repos repos ON dws.repo_id = repos.id
WHERE MONTH >= '2015-01' AND MONTH <='2022-12'
GROUP BY dws.repo_id,
 repos.name
ORDER BY stars DESC LIMIT 10;