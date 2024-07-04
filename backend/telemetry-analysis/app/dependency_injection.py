from dependency_injector import containers, providers
from infraestructure.server.controllers.measures import MeasuresIoTServicer
from infraestructure.server.server import GRPCServer

from app.application import Application
class AppContainer(containers.DeclarativeContainer):
    config = providers.Configuration()


    #repositories

    #usecases

    #controllers
    controller_measures = providers.Factory(MeasuresIoTServicer)
    #infraestructure
    grpc_server = providers.Singleton(GRPCServer, measuresServicer=controller_measures)
    #application
    app = providers.Factory(Application, server=grpc_server)