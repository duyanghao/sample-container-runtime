package subsystems

type ResourceConfig struct {
	MemoryLimit string `json:"memoryLimit"`
	CpuShare    string `json:"cpuShare"`
	CpuSet      string `json:"cpuSet"`
}

type Subsystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
