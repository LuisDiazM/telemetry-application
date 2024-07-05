from dependency_injector import containers, providers
from domain.measures.usecases.usecase import MeasuresUsecase
from infraestructure.database.mongo_imp import MongoImp
from infraestructure.database.repositories.measures.measures_repo import MeasuresDBRepository
from infraestructure.server.controllers.measures import MeasuresIoTServicer
from infraestructure.server.server import GRPCServer

from app.application import Application
class AppContainer(containers.DeclarativeContainer):
    config = providers.Configuration()

    mongo = providers.Singleton(MongoImp)

    #repositories
    db_repo = providers.Factory(MeasuresDBRepository, mongo=mongo)
    #usecases
    measures_usecase = providers.Factory(MeasuresUsecase, measures_db=db_repo)

    #controllers
    controller_measures = providers.Factory(MeasuresIoTServicer, measures_usecase=measures_usecase)
    #infraestructure
    grpc_server = providers.Singleton(GRPCServer, measuresServicer=controller_measures)
    #application
    app = providers.Factory(Application, server=grpc_server)