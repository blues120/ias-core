package data

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"github.com/blues120/ias-kit/oss"
	"github.com/blues120/ias-kit/oss/local"
	"github.com/blues120/ias-kit/oss/s3"
)

func NewOssRepo(c *conf.Data) (oss.Oss, error) {
	switch cfg := c.Oss.Oss.(type) {
	case *conf.Data_Oss_Local_:
		return local.NewLocal(cfg.Local.StorePath, cfg.Local.Path)
	case *conf.Data_Oss_AwsS3_:
		repo, err := s3.NewS3(cfg.AwsS3.Bucket,
			&aws.Config{
				Credentials:      credentials.NewStaticCredentials(cfg.AwsS3.Ak, cfg.AwsS3.Sk, ""),
				Endpoint:         aws.String(cfg.AwsS3.Endpoint),
				S3ForcePathStyle: aws.Bool(true),
				DisableSSL:       aws.Bool(true),
				Region:           aws.String(cfg.AwsS3.Region),
			}, cfg.AwsS3.EndpointAlias)

		return repo, err
	default:
		return nil, fmt.Errorf("unsupported oss mode: %s", cfg)
	}
}
