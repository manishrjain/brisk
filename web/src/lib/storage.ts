const STORAGE_KEY = 'rentobuy_saved_configs';

export interface SavedConfig {
  name: string;
  data: Record<string, any>;
  savedAt: number;
}

export function getSavedConfigs(): SavedConfig[] {
  if (typeof window === 'undefined') return [];

  const saved = localStorage.getItem(STORAGE_KEY);
  if (!saved) return [];

  try {
    return JSON.parse(saved);
  } catch {
    return [];
  }
}

export function getSavedConfigNames(): string[] {
  return getSavedConfigs().map(c => c.name);
}

export function saveConfig(name: string, data: Record<string, any>): void {
  const configs = getSavedConfigs();
  const existingIndex = configs.findIndex(c => c.name === name);

  const newConfig: SavedConfig = {
    name,
    data,
    savedAt: Date.now(),
  };

  if (existingIndex >= 0) {
    configs[existingIndex] = newConfig;
  } else {
    configs.push(newConfig);
  }

  localStorage.setItem(STORAGE_KEY, JSON.stringify(configs));
}

export function loadConfig(name: string): Record<string, any> | null {
  const configs = getSavedConfigs();
  const config = configs.find(c => c.name === name);
  return config?.data ?? null;
}

export function deleteConfig(name: string): void {
  const configs = getSavedConfigs();
  const filtered = configs.filter(c => c.name !== name);
  localStorage.setItem(STORAGE_KEY, JSON.stringify(filtered));
}
