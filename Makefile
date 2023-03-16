run:
	nodemon --exec go run . --signal SIGTERM
update_bot_image:
	docker build -t personal_bot:$(VERSION) .
update_bot_db_image:
	docker build -t personal_bot_db:$(VERSION) ./db/.

.PHONY: run update_bot_image