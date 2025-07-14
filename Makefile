lint-breaking::
	go tool buf breaking --against 'https://github.com/LamichhaneBibek/news-grpc.git#branch=main'

# Lint the proto files
lint-proto::
	go tool buf lint --config buf.yaml

# Generate the go code from the proto files
generate-proto::
	go tool buf generate --template buf.gen.yaml
