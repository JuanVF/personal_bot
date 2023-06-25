#Copyright 2023 Juan Jose Vargas Fletes
#
#This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
#To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/
#
#Under the CC BY-NC license, you are free to:
#
#- Share: copy and redistribute the material in any medium or format
#- Adapt: remix, transform, and build upon the material
#
#Under the following terms:
#
#  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
#    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.
#
#- Non-Commercial: You may not use the material for commercial purposes.
#
#You are free to use this work for personal or non-commercial purposes.
#If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
run:
	nodemon --exec go run . --signal SIGTERM
update_bot_image:
	docker build -t personal_bot:$(VERSION) .
update_bot_db_image:
	docker build -t personal_bot_db:$(VERSION) ./db/.
run_test:
	set ENVIRONMENT=test&& go test ./... -v
format:
	gofmt -w -s .
lint:
	gofmt -l .

# This will create a new SQL File for a migration with the correct copyright notice. Just pass the migration name as parameter
new_migration:
	@touch ./db/migrations/$$(date "+%Y%m%d%H%M")-$(NAME).sql
	@echo "/*Copyright 2023 Juan Jose Vargas Fletes\n"  \
			"This work is licensed under the Creative Commons Attribution-NonCommercial \(CC BY-NC\) license.\n" \
			"To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/\n\n" \
			"Under the CC BY-NC license, you are free to:\n\n" \
			"- Share: copy and redistribute the material in any medium or format\n" \
			"- Adapt: remix, transform, and build upon the material\n\n" \
			"Under the following terms:\n\n" \
			"- Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.\n" \
			"	You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.\n\n" \
			"- Non-Commercial: You may not use the material for commercial purposes.\n\n" \
			"You are free to use this work for personal or non-commercial purposes.\n" \
			"If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.\n" \
			"*/" >> ./db/migrations/$$(date "+%Y%m%d%H%M")-$(NAME).sql
	@echo "File [./db/migrations/$$(date "+%Y%m%d%H%M")-$(NAME).sql] Created"

.PHONY: run update_bot_image
.PHONY: run update_bot_db_image