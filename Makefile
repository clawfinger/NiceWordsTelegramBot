build:
	go build

docker:
	sudo docker build --build-arg BOT_TOKEN=7364207728:AAHYwpmonWtthagDreJ4toNymSRUUtlqGak --build-arg CHANNEL_ID=-4223240467 --build-arg=https://pastebin.com/raw/V9khBcPb .