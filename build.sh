# mkdir pkgs 
cp go.mod ./bitbucket_client
cd bitbucket_client
go mod tidy
# -v
# 	print the names of packages as they are compiled.
# -work
# 	print the name of the temporary work directory and
# 	do not delete it when exiting.
# -x
# 	print the commands. 
go build -v -work -x -a ./...
cd ..
