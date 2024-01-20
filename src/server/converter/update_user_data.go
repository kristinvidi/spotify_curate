package converter

import pb "src/server/proto"

func SerializeUpdateUserData(success bool) *pb.UpdateUserDataResponse {
	return &pb.UpdateUserDataResponse{
		General: SerializeGeneral(success),
	}
}
