@echo protocol generate
@echo current dir:%~dp0
set cDir= %~dp0


protoc  --go_out=./gen  proto/*.proto