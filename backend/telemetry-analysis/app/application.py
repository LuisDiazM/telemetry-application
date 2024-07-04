from infraestructure.server.server import GRPCServer
class Application:
    def __init__(self, server:GRPCServer) -> None:
        self.server = server

    def start(self):
        self.server.registry_controllers()
        self.server.start_server()
    