from concurrent import futures
import grpc
import os
from infraestructure.server.protos.measures import measures_pb2_grpc
from infraestructure.server.controllers.measures import MeasuresIoTServicer

class GRPCServer:
    def __init__(self, measuresServicer:MeasuresIoTServicer) -> None:
        workers = int(os.getenv("WORKERS",4))
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=workers))
        self.measuresServicer = measuresServicer

    def registry_controllers(self):
        # registry all controllers grpc
        measures_pb2_grpc.add_MeasuresIoTServicer_to_server(self.measuresServicer, self.server)

    def start_server(self):
        PORT = os.getenv("PORT",50001)
        self.server.add_insecure_port(f"[::]:{PORT}")
        self.server.start()
        self.server.wait_for_termination()