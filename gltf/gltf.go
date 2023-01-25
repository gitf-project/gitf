package gltf

type GLTF struct {
	Raw string
}

func New(raw string) (gltf *GLTF) {
	gltf.Raw = raw
	return gltf
}

// 获取场景数量
func (gltf *GLTF) GetScenesNumber() {}

func (gltf *GLTF) GetNodesNumber() {}

func (gltf *GLTF) GetCamerasNumber() {}

func (gltf *GLTF) GetAnimationsNumber() {}

func (gltf *GLTF) GetImagesNumber() {}

func (gltf *GLTF) GetBuffersNumber() {}

func (gltf *GLTF) GetMeshesNumber() {}

func (gltf *GLTF) GetBufferViewsNumber() {}

func (gltf *GLTF) GetAccessorsNumber() {}
