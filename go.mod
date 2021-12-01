module rickover

go 1.16

replace metrics v0.0.0 => ./go-simple-metrics

require (
	github.com/Shyp/go-dberror v0.0.0-20160419232342-92eb70dfc808
	github.com/Shyp/go-debug v2.0.1-0.20160414204915-8a83ccc090f8+incompatible // indirect
	github.com/Shyp/go-types v0.0.0-20160531142844-d3a92d5d0264
	github.com/Shyp/rest v0.0.0-20171024210032-b2c053a6ab95
	github.com/gorilla/handlers v0.0.0-20150412184004-a24b39a6a2c8
	github.com/inconshreveable/log15 v0.0.0-20171019012758-0decfc6c20d9 // indirect
	github.com/kevinburke/rest v0.0.0-20210506044642-5611499aa33c // indirect
	github.com/letsencrypt/boulder v0.0.0-20211129192819-2a2629493dd7 // indirect
	github.com/lib/pq v1.10.4
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	metrics v0.0.0
)
