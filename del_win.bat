@echo off
cd %~dp0
for /D %%s in (*) do ( 
cd %%s
del /s /q *.exe
if exist go.mod (
del go.mod
)
cd .. 
)