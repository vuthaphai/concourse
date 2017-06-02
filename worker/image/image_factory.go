package image

import (
	"errors"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

var ErrUnsupportedResourceType = errors.New("unsupported resource type")

type imageFactory struct {
	imageResourceFetcherFactory ImageResourceFetcherFactory
}

func NewImageFactory(
	imageResourceFetcherFactory ImageResourceFetcherFactory,
) worker.ImageFactory {
	return &imageFactory{
		imageResourceFetcherFactory: imageResourceFetcherFactory,
	}
}

func (f *imageFactory) GetImage(
	logger lager.Logger,
	worker worker.Worker,
	volumeClient worker.VolumeClient,
	imageSpec worker.ImageSpec,
	teamID int,
	cancel <-chan os.Signal,
	delegate worker.ImageFetchingDelegate,
	resourceUser db.ResourceUser,
	resourceTypes atc.VersionedResourceTypes,
) (worker.Image, error) {
	if imageSpec.ImageArtifactSource != nil {
		artifactVolume, existsOnWorker, err := imageSpec.ImageArtifactSource.VolumeOn(worker)
		if err != nil {
			logger.Error("failed-to-check-if-volume-exists-on-worker", err)
			return nil, err
		}

		if existsOnWorker {
			return &imageProvidedByPreviousStepOnSameWorker{
				artifactVolume: artifactVolume,
				imageSpec:      imageSpec,
				teamID:         teamID,
				volumeClient:   volumeClient,
			}, nil
		}

		return &imageProvidedByPreviousStepOnDifferentWorker{
			imageSpec:    imageSpec,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	// check if custom resource type
	for _, resourceType := range resourceTypes {
		if resourceType.Name == imageSpec.ResourceType {
			imageResourceFetcher := f.imageResourceFetcherFactory.ImageResourceFetcherFor(worker)
			imageParentVolume, imageMetadataReader, version, err := imageResourceFetcher.Fetch(
				logger.Session("image"),
				cancel,
				resourceUser,
				resourceType.Type,
				resourceType.Source,
				worker.Tags(),
				teamID,
				resourceTypes.Without(imageSpec.ResourceType),
				delegate,
				resourceType.Privileged,
			)
			if err != nil {
				logger.Error("failed-to-fetch-image", err)
				return nil, err
			}

			return &imageFromResource{
				imageParentVolume:   imageParentVolume,
				version:             version,
				imageMetadataReader: imageMetadataReader,
				privileged:          resourceType.Privileged,
				teamID:              teamID,
				volumeClient:        volumeClient,
			}, nil
		}
	}

	if imageSpec.ImageResource != nil {
		imageResourceFetcher := f.imageResourceFetcherFactory.ImageResourceFetcherFor(worker)
		imageParentVolume, imageMetadataReader, version, err := imageResourceFetcher.Fetch(
			logger.Session("image"),
			cancel,
			resourceUser,
			imageSpec.ImageResource.Type,
			imageSpec.ImageResource.Source,
			worker.Tags(),
			teamID,
			resourceTypes,
			delegate,
			imageSpec.Privileged,
		)
		if err != nil {
			logger.Error("failed-to-fetch-image", err)
			return nil, err
		}

		return &imageFromResource{
			imageParentVolume:   imageParentVolume,
			version:             version,
			imageMetadataReader: imageMetadataReader,
			privileged:          imageSpec.Privileged,
			teamID:              teamID,
			volumeClient:        volumeClient,
		}, nil
	}

	if imageSpec.ResourceType != "" {
		return &imageFromBaseResourceType{
			worker:           worker,
			resourceTypeName: imageSpec.ResourceType,
			teamID:           teamID,
			volumeClient:     volumeClient,
		}, nil
	}

	return &imageFromRootfsURI{
		url: imageSpec.ImageURL,
	}, nil
}
