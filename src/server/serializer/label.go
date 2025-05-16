package serializer

import pb "src/server/proto"

func SerializeCreateLabelsForUserResponse(success bool, failedLabels []string) *pb.CreateLabelsForUserResponse {
	var failureDetails *string
	if len(failedLabels) > 0 {
		details := "some labels could not be created"
		failureDetails = &details
	}

	return &pb.CreateLabelsForUserResponse{
		General:      SerializeGeneral(success, failureDetails),
		FailedLabels: failedLabels,
	}
}
