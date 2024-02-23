CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/auctionBot cmd/bot/main.go
zip -r bin/auctionBot_linux_x64.zip bin/auctionBot
CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -o bin/auctionBot cmd/bot/main.go
zip -r bin/auctionBot_linux_x86.zip bin/auctionBot
CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/auctionBot cmd/bot/main.go
zip -r bin/auctionBot_osx_x64.zip bin/auctionBot
CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/auctionBot cmd/bot/main.go
zip -r bin/auctionBot_osx_arm.zip bin/auctionBot
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o bin/auctionBot.exe cmd/bot/main.go
zip -r bin/auctionBot_win_x64.zip bin/auctionBot.exe
CGO_ENABLED=0 GOARCH=386 GOOS=windows go build -o bin/auctionBot.exe cmd/bot/main.go
zip -r bin/auctionBot_win_x86.zip bin/auctionBot.exe
