package brimstoneesan

const version = "1.0.0"

type Brimstoneesan struct {
	AppName string
	Debug   bool
	Version string
}

func (b *Brimstoneesan) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := b.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (b *Brimstoneesan) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := b.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
