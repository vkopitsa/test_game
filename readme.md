go get -u github.com/golang/protobuf/protoc-gen-go

cd messages/ 
protoc --go_out=. --js_out=library=../public/pb,binary:. *.proto
