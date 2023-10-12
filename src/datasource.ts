import { DataSourceInstanceSettings, CoreApp } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';

import { MaxComputeQuery, MaxcomputeDataSourceOptions, DEFAULT_QUERY } from './types';

export class DataSource extends DataSourceWithBackend<MaxComputeQuery, MaxcomputeDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<MaxcomputeDataSourceOptions>) {
    super(instanceSettings);
  }

  getDefaultQuery(_: CoreApp): Partial<MaxComputeQuery> {
    return DEFAULT_QUERY
  }
}
