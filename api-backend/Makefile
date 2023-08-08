.PHONY: seed

seed: 
	docker exec -it deployments-postgres-1 psql -U postgres -d testDB -f ./go/scripts/seed_db.sql