mocks:
	mockery --all --keeptree

run:
	$(MAKE) infra-up && \
		go run .

infra-up:	
	$(MAKE) infra-down && \
	docker-compose -f docker-compose.yml up -d && \
		docker exec userbirthday-mysql sh -c 'chmod +x /script/wait-mysql.sh; /script/wait-mysql.sh' && \
		$(MAKE) migrate-up

infra-down:
	docker-compose -f docker-compose.yml down

migrate-up:
	docker run --env-file .env \
		--network=host \
		-v $(PWD)/infrastructure/repository/mysql/migration:/migration \
		amacneil/dbmate up

migrate-down:
	docker run --env-file .env \
		--network=host \
		-v $(PWD)/infrastructure/repository/mysql/migration:/migration \
		amacneil/dbmate down

unit-test:
	go test ./service/... \
		-failfast -v

integration-test:
	$(MAKE) test-infra-up && \
	go test ./infrastructure/repository/... \
		-failfast -v

all-test:
	$(MAKE) unit-test && \
		$(MAKE) integration-test

test-infra-up:	
	$(MAKE) test-infra-down && \
	docker-compose -f docker-compose.test.yml up -d && \
		docker exec userbirthday-mysql-test sh -c 'chmod +x /script/wait-mysql.sh; /script/wait-mysql.sh' && \
		$(MAKE) test-migrate-up

test-infra-down:
	docker-compose -f docker-compose.test.yml down

test-migrate-up:
	docker run --env-file .env.testing \
		--network=host \
		-v $(PWD)/infrastructure/repository/mysql/migration:/migration \
		amacneil/dbmate up

test-migrate-down:
	docker run --env-file .env.testing \
		--network=host \
		-v $(PWD)/infrastructure/repository/mysql/migration:/migration \
		amacneil/dbmate down
	
.PHONY: mocks