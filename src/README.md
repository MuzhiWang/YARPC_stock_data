`cd .\idl`
`protoc.exe --descriptor_set_out=. --gogoslick_out=. --yarpc-go_out=. .\stock_data.proto`



`.\yab.exe -p localhost:5432 --service stockdata --procedure  muzwang.stock_data.StockData/Test -F C:\Users\wmz66\go\src\StockData\src\idl\stock_data.pb.h --timeout 90000 -f .\test-request.json`