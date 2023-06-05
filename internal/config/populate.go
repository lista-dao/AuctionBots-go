package config

type populate func()

func (cfg *config) populate() {
	populates := []populate{
		cfg.analyticsClient.populate,
		cfg.ethereum.populate,
		cfg.contracts.populate,
		cfg.logger.populate,
		cfg.settings.populate,
	}

	for _, populate := range populates {
		populate()
	}
}
