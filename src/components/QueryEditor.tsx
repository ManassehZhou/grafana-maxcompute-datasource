import React, { useState } from 'react';
import { CodeEditor, InlineField, InlineFieldRow, Select, InlineSwitch, QueryField } from '@grafana/ui';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from '../datasource';
import { MaxcomputeDataSourceOptions, MaxComputeQuery } from '../types';

type Props = QueryEditorProps<DataSource, MaxComputeQuery, MaxcomputeDataSourceOptions>;

function calculateHeight(queryText: string): number {
  const minHeight = 200;
  const maxHeight = 500;

  // assume 20 px per row
  let desiredHeight = queryText.split('\n').length * 20;

  // return the value in a range between the min and max height
  return Math.min(maxHeight, Math.max(minHeight, desiredHeight));
}

export function QueryEditor({ query, onChange, onRunQuery }: Props) {
  const [useLegacyEditor, setUseLegacyEditor] = useState(false);

  const queryTypes: Array<SelectableValue<string>> = [
    {
      label: 'Timeseries',
      value: 'timeseries',
    },
    {
      label: 'Table',
      value: 'table',
    },
  ];
  const queryType = queryTypes.find((type) => type.value === query.queryType) || queryTypes[1];

  const onChangeQueryType = (value: SelectableValue<string>) => {
    console.log(value.value)
    onChange({
      ...query,
      queryType: value.value || 'table',
    });
    onRunQuery();
  };

  const onChangeRawQuery = (rawQueryText: string) => {
    onChange({
      ...query,
      rawQueryText: rawQueryText,
    });
    onRunQuery();
  };

  return (
    <>
      <div>
        {useLegacyEditor ? (
          <QueryField
            query={query.rawQueryText}
            onBlur={() => onRunQuery()}
            onChange={onChangeRawQuery}
            portalOrigin='mock-origin'
          />
        ): (
          <CodeEditor
            height={calculateHeight(query.rawQueryText)}
            value={query.rawQueryText}
            onBlur={onChangeRawQuery}
            onSave={onChangeRawQuery}
            language='sql'
          />
        )}
      </div>
      <InlineFieldRow>
        <InlineField label='Query type'>
          <Select options={queryTypes} onChange={onChangeQueryType} value={queryType} />
        </InlineField>

        <InlineField label="Use legacy editor">
          <InlineSwitch value={useLegacyEditor} onChange={() => setUseLegacyEditor(!useLegacyEditor)} />
        </InlineField>
      </InlineFieldRow>
    </>
  );
}
