go get -u github.com/golang/protobuf/protoc-gen-go

cd messages/ && protoc --go_out=. *.proto
protoc --js_out=library=../public/pb,binary:. *.proto

pbjs -t json file1.proto file2.proto > bundle.json