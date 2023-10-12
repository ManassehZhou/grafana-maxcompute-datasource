import { DataSourceInstanceSettings, CoreApp } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';

import { MCQuery, MCConfig, defaultMCSQLQuery } from './types';

export class DataSource extends DataSourceWithBackend<MCQuery, MCConfig> {
  constructor(instanceSettings: DataSourceInstanceSettings<MCConfig>) {
    super(instanceSettings);
  }

  getDefaultQuery(_: CoreApp): Partial<MCQuery> {
    return defaultMCSQLQuery; 
  }
}
