# Build the latest version of the project
go build
# Replace the executable with the new one
cd ./bin
rm ./nebula
cd ..
mv nebula bin
# Run the executable, passing in any arguments
./bin/nebula $*
