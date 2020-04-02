@echo off
cd %~dp0
call del_win.bat
for /D %%s in (*) do ( 
cd %%s
go mod init "%%s"
go build
cd .. 
)
