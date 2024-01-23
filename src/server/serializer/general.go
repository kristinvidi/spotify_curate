package serializer

import pb "src/server/proto"

func SerializeGeneral(success bool, failureDetails *string) *pb.General {
	var pbFailureDetails string
	if failureDetails != nil {
		pbFailureDetails = *failureDetails
	}

	return &pb.General{Success: success, FailureDetails: pbFailureDetails}
}
