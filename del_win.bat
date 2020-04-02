@echo off
cd %~dp0
for /D %%s in (*) do ( 
cd %%s
echo %%s
del /s /q *.exe
if exist go.mod (
echo del go.mod
del go.mod
)
cd .. 
)