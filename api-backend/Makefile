.PHONY: seed

seed: 
	docker exec -it deployments-postgres-1 psql -U postgres -d api-backend -f ./go/scripts/seed_db.sql