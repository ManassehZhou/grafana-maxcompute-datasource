import { DataQuery, DataSourceJsonData } from '@grafana/data';

export interface MaxComputeQuery extends DataQuery {
  rawQueryText: string;
  queryType: string;
}

export const DEFAULT_QUERY: Partial<MaxComputeQuery> = {
  rawQueryText: `SELECT 1;`,
  queryType: 'table',
};

/**
 * These are options configured for each DataSource instance
 */
export interface MaxcomputeDataSourceOptions extends DataSourceJsonData {
  accessKeyId?: string;
  endpoint?: string;
  project?: string;
}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface MaxcomputeSecureJsonData {
  accessKeySecret?: string;
}
