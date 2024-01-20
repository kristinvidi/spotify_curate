package converter

import pb "src/server/proto"

func SerializeGeneral(success bool) *pb.General {
	return &pb.General{Success: success}
}
