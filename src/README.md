`cd .\idl`
`protoc.exe --gogoslick_out=. --yarpc-go_out=. .\stock_data.proto`



`.\yab.exe  -p localhost:5432 --procedure muzwang.stock_data.StockData/Test -r '{"value":"ttt"}' --service stock-data`