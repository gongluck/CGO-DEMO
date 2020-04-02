@echo off
cd %~dp0
call del_win.bat
for /D %%s in (*) do ( 
echo %%s
cd %%s
go mod init "%%s"
go build
cd .. 
)
