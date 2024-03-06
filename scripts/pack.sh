CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o auctionBot/auctionBot cmd/bot/main.go
zip -r auctionBot/auctionBot_linux_x64.zip auctionBot/auctionBot auctionBot/config.txt auctionBot/config/config.yaml
CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -o auctionBot/auctionBot cmd/bot/main.go
zip -r auctionBot/auctionBot_linux_x86.zip auctionBot/auctionBot auctionBot/config.txt auctionBot/config/config.yaml
CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o auctionBot/auctionBot cmd/bot/main.go
zip -r auctionBot/auctionBot_osx_x64.zip auctionBot/auctionBot auctionBot/config.txt auctionBot/config/config.yaml
CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o auctionBot/auctionBot cmd/bot/main.go
zip -r auctionBot/auctionBot_osx_arm.zip auctionBot/auctionBot auctionBot/config.txt auctionBot/config/config.yaml
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o auctionBot/auctionBot.exe cmd/bot/main.go
zip -r auctionBot/auctionBot_win_x64.zip auctionBot/auctionBot.exe auctionBot/config.txt auctionBot/config/config.yaml
CGO_ENABLED=0 GOARCH=386 GOOS=windows go build -o auctionBot/auctionBot.exe cmd/bot/main.go
zip -r auctionBot/auctionBot_win_x86.zip auctionBot/auctionBot.exe auctionBot/config.txt auctionBot/config/config.yaml
