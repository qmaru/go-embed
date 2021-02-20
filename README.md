# go-embed
Pack [gin-gonic](https://github.com/gin-gonic/gin) and [create-react-app](https://github.com/facebook/create-react-app) into a binary file.  

## Usage
+ go1.16+
+ create-react-app

After running the file, the system browser will be used to access the monitoring address.
```go
cd web
yarn install
yarn build
cd ..
go build -o server main.go
./server 
```