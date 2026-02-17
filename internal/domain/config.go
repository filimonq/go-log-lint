package domain

type Config struct {
	Rules map[string]RuleConfig `yaml:"rules"`
}

type RuleConfig struct {
	Enabled  bool     `yaml:"enabled"`
	Patterns []string `yaml:"patterns,omitempty"`
}

func DefaultConfig() *Config {
	return &Config{
		Rules: map[string]RuleConfig{
			"LowercaseRule":     {Enabled: true},
			"SpecialCharsRule":  {Enabled: true},
			"SensitiveDataRule": {Enabled: true},
			"EnglishOnlyRule":   {Enabled: true},
		},
	}
}
