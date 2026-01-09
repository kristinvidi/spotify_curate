.PHONY: run-backend run-frontend run proto stop

stop:
	@echo "Stopping Frontend (Port 3000)..."
	@-lsof -ti:3000 | xargs kill -9 2>/dev/null || true
	@echo "Stopping Backend (Port 50051)..."
	@-lsof -ti:50051 | xargs kill -9 2>/dev/null || true
	@echo "Services stopped."

proto:
	@echo "Generating Go Proto files..."
	@PATH=$$PATH:$$(go env GOPATH)/bin protoc --proto_path=proto --go_out=src/server/proto --go_opt=paths=source_relative --go-grpc_out=src/server/proto --go-grpc_opt=paths=source_relative proto/spotify_curate.proto
	@echo "Copying Proto file to Frontend..."
	@mkdir -p web/proto
	@cp proto/spotify_curate.proto web/proto/spotify_curate.proto
	@echo "Proto generation complete."


run-backend:
	@echo "Starting Backend..."
	cd src && go run main.go

run-frontend:
	@echo "Starting Frontend..."
	cd web && npm run dev

run:
	@echo "Starting Spotify Curator App..."
	@(trap 'kill 0' SIGINT; \
		$(MAKE) run-backend & \
		$(MAKE) run-frontend & \
		wait)
