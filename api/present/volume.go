package present

import (
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

func Volume(volume dbng.CreatedVolume) atc.Volume {
	return atc.Volume{
		ID:              volume.Handle(),
		Type:            string(volume.Type()),
		WorkerName:      volume.Worker().Name,
		SizeInBytes:     volume.SizeInBytes(),
		ContainerHandle: volume.ContainerHandle(),
		ParentHandle:    volume.ParentHandle(),
	}
}
