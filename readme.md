go get -u github.com/golang/protobuf/protoc-gen-go
npm install --save-dev ts-protoc-gen

cd messages/ 
protoc --go_out=. --js_out=library=../public/pb,binary:. *.proto

or 

./generate.sh
