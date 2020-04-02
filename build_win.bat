@echo off
cd %~dp0
for /D %%s in (*) do ( 
echo %%s
cd %%s
go mod init test
go build
cd .. 
)
