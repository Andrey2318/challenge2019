package v1

import (
	"challenge2019/internal/application/traffic"
	"challenge2019/internal/config"
	"challenge2019/internal/domain/output"
	"challenge2019/internal/infrastructure/services/parser/csv"
	"github.com/urfave/cli/v2"
	"strconv"
)

type ChallengeService struct {
	trafficApp *traffic.UseCase
}

func New(trafficApp *traffic.UseCase) *ChallengeService {
	return &ChallengeService{trafficApp: trafficApp}
}

func (s *ChallengeService) Statement1() *cli.Command {
	return &cli.Command{
		Name:  "statement1",
		Usage: "",
		Flags: []cli.Flag{&cli.StringFlag{
			Name:  "input",
			Value: "",
		}},
		Action: func(ctx *cli.Context) error {
			fileName := ctx.String("input")
			if fileName == "" {
				fileName = config.GetInputFile()
			}

			idata, err := csv.ParseInput(fileName)
			if err != nil {
				return err
			}

			input := make([]*traffic.StatementRequest, len(idata))
			for i := range idata {
				input[i] = &traffic.StatementRequest{
					DeliveryID: idata[i].DeliveryID,
					Size:       idata[i].Size,
					Theatre:    idata[i].Theatre,
				}
			}

			data, err := s.trafficApp.Statement1(ctx.Context, input)
			if err != nil {
				return err
			}

			Output := make([]*output.Model, len(data))
			for i, datum := range data {
				Output[i] = &output.Model{
					DeliveryID: datum.DeliveryID,
					Status:     datum.Status,
					Partner:    datum.PartnerID,
				}
				if datum.Cost > 0 {
					Output[i].Cost = strconv.FormatUint(uint64(datum.Cost), 10)
				}
			}

			return csv.SaveOutput("output1.csv", Output)
		},
	}
}

func (s *ChallengeService) Statement2() *cli.Command {
	return &cli.Command{
		Name:  "statement2",
		Usage: "",
		Flags: []cli.Flag{&cli.StringFlag{
			Name:  "input",
			Value: "",
		}},
		Action: func(ctx *cli.Context) error {
			fileName := ctx.String("input")
			if fileName == "" {
				fileName = config.GetInputFile()
			}

			idata, err := csv.ParseInput(fileName)
			if err != nil {
				return err
			}

			input := make([]*traffic.StatementRequest, len(idata))
			for i := range idata {
				input[i] = &traffic.StatementRequest{
					DeliveryID: idata[i].DeliveryID,
					Size:       idata[i].Size,
					Theatre:    idata[i].Theatre,
				}
			}

			data, err := s.trafficApp.Statement2(ctx.Context, input)
			if err != nil {
				return err
			}

			Output := make([]*output.Model, len(data))
			for i, datum := range data {
				Output[i] = &output.Model{
					DeliveryID: datum.DeliveryID,
					Status:     datum.Status,
					Partner:    datum.PartnerID,
				}
				if datum.Cost > 0 {
					Output[i].Cost = strconv.FormatUint(uint64(datum.Cost), 10)
				}
			}

			return csv.SaveOutput("output2.csv", Output)
		},
	}
}
