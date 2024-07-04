.PHONY = pb2

SHELL := /bin/bash

pb2:
	python3 -m venv env
	source env/bin/activate
	pip3 install grpcio grpcio-tools
	
	python -m grpc_tools.protoc -I./infraestructure/server/protos/measures --python_out=./infraestructure/server/protos/measures --pyi_out=./infraestructure/server/protos/measures --grpc_python_out=./infraestructure/server/protos/measures ./infraestructure/server/protos/measures/measures.proto
	rm -r env
