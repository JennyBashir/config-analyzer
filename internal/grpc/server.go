package grpc

import (
	"context"
	"errors"

	analyzer "github.com/JennyBashir/config-analyzer/gen"
	issues "github.com/JennyBashir/config-analyzer/internal/analyzer"
	"github.com/JennyBashir/config-analyzer/internal/config"
)

type Server struct {
	analyzer.UnimplementedConfigAnalyzerServer
}

func (s *Server) Analyze(ctx context.Context, req *analyzer.AnalyzeRequest) (*analyzer.AnalyzeResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	data := []byte(req.Config)

	var cfg config.Config
	var err error

	extSlice := []string{
		".json",
		".yaml",
		".yml",
	}

	for _, ext := range extSlice {
		cfg, err = config.Parse(data, ext)
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	iss := issues.Analyze(cfg)
	var res []*analyzer.Issue

	for i := 0; i < len(iss); i++ {
		res = append(res, &analyzer.Issue{
			Severity:       iss[i].Severity,
			Message:        iss[i].Message,
			Recommendation: iss[i].Recommendation,
		})
	}

	return &analyzer.AnalyzeResponse{
		Issues: res,
	}, nil
}
