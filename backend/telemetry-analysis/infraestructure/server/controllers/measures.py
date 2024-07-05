from infraestructure.server.protos.measures import measures_pb2_grpc, measures_pb2
from infraestructure.server.controllers.utils.adapter_responses import adapter_measures_response
from domain.measures.usecases.usecase import MeasuresUsecase
import grpc
class MeasuresIoTServicer(measures_pb2_grpc.MeasuresIoTServicer):
    def __init__(self, measures_usecase:MeasuresUsecase) -> None:
        self.measures_usecase = measures_usecase
        super().__init__()
    
    def GetMeasuresBySensor(self, request, context):
        sensor_id = request.device_id
        data = self.measures_usecase.get_measures_by_device(sensor_id)
        response = adapter_measures_response(data)
        if len(response.values)>0:
            return response
        else:
            context.abort(grpc.StatusCode.NOT_FOUND,"no data")