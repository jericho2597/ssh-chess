COMPOSE=docker-compose -f docker-compose.yml

CHESS_ENGINE=chess-engine
SSH_SERVER=ssh-server

# Build all services
.PHONY: build
build:
	$(COMPOSE) build

# start services and listen to logs
.PHONY: start
start:
	$(COMPOSE) up

# start services detached
.PHONY: start-detached
start-detached:
	$(COMPOSE) up -d

# Build all services
.PHONY: stop
stop:
	$(COMPOSE) down


# hot-reloads just the chess-engine container
.PHONY: hot-reload-chess-engine
hot-reload-chess-engine:
	$(COMPOSE) build $(CHESS_ENGINE) && $(COMPOSE) up -d --no-deps $(CHESS_ENGINE)

# hot-reloads just the ssh-server container
.PHONY: hot-reload-ssh-server
hot-reload-ssh-server:
	$(COMPOSE) build $(SSH_SERVER) && $(COMPOSE) up -d --no-deps $(SSH_SERVER)

# takes down all services and cleans artifacts
.PHONY: clean
clean:
	$(COMPOSE) down --volumes --remove-orphans

.PHONY: status
status:
	$(COMPOSE) ps