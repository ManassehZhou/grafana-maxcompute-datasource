import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { ConfigEditor } from './components/ConfigEditor';
import { QueryEditor } from './components/QueryEditor';
import { MaxComputeQuery, MaxcomputeDataSourceOptions } from './types';

export const plugin = new DataSourcePlugin<DataSource, MaxComputeQuery, MaxcomputeDataSourceOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
