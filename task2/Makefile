.PHONY: all

start-servers:
	@echo "  >  Building binary..."
	sudo docker-compose up -d postgres

init-postgresql: start-servers
	sudo docker-compose exec -T postgres psql -U user -d dbname < ddl/create_db.sql

clean:
	sudo docker-compose down --volumes --remove-orphans