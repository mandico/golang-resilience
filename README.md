### Run Docker MSSQL
docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=P@ssw0rd1234" \
   -p 1433:1433 --name sql --hostname sql \
   -d \
   mcr.microsoft.com/mssql/server:2022-latest

   docker exec -t sql cat /var/opt/mssql/log/errorlog | grep connection

   docker exec -it sql "bash"

   /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "P@ssw0rd1234"

SELECT @@SERVERNAME,SERVERPROPERTY('ComputerNamePhysicalNetBIOS'),SERVERPROPERTY('MachineName'),SERVERPROPERTY('ServerName');