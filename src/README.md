`cd .\idl`
`protoc.exe --descriptor_set_out=stock_data.pb.h --gogoslick_out=. --yarpc-go_out=. .\stock_data.proto`

`cd .\idl\tdx`
`protoc.exe --gogoslick_out=. --yarpc-go_out=. .\TdxReader.proto`



`.\yab.exe -p localhost:5432 --service stockdata --procedure  muzwang.stock_data.StockData/Test -F C:\Users\wmz66\go\src\StockData\src\idl\stock_data.pb.h --timeout 90000 -f .\test-request.json`


`yab -p localhost:5432 --service stockdata --procedure  muzwang.stock_data.StockData/Test -F /Users/superegg/go/src/YARPC_stock_data/src/idl/stock_data.pb.h -r '{"value": "from yabbbbbb"}'`
