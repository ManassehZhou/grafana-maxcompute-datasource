import React, { ChangeEvent } from 'react';
import { InlineField, Input, SecretInput } from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { MaxcomputeDataSourceOptions, MaxcomputeSecureJsonData } from '../types';

interface Props extends DataSourcePluginOptionsEditorProps<MaxcomputeDataSourceOptions> {}

export function ConfigEditor(props: Props) {
  const { onOptionsChange, options } = props;
  const onAccessKeyIdChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      accessKeyId: event.target.value,
    }
    onOptionsChange({ ... options, jsonData });
  };

  const onAccessKeySecretChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({ 
      ... options, 
      secureJsonData: {
        accessKeySecret: event.target.value,
      },
    });
  };

  const onResetAccessKeySecret = () => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        accessKeySecret: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        accessKeySecret: '',
      },
    });
  };

  const onEndpointChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      endpoint: event.target.value,
    }
    onOptionsChange({ ... options, jsonData });
  };

  const onProjectChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      project: event.target.value,
    }
    onOptionsChange({ ... options, jsonData });
  };

  const { jsonData, secureJsonFields } = options;
  const secureJsonData = (options.secureJsonData || {}) as MaxcomputeSecureJsonData;

  return (
    <div className="gf-form-group">
      <InlineField label="AccessKey ID" labelWidth={18}>
        <Input
          onChange={onAccessKeyIdChange}
          value={jsonData.accessKeyId || ''}
          placeholder='Aliyun Access Key ID'
          width={40}
        />
      </InlineField>

      <InlineField label="AccessKey Secret" labelWidth={18}>
        <SecretInput
          isConfigured={(secureJsonFields && secureJsonFields.accessKeySecret) as boolean}
          value={secureJsonData.accessKeySecret}
          placeholder='Aliyun Access Key Secret'
          width={40}
          onReset={onResetAccessKeySecret}
          onChange={onAccessKeySecretChange}
        />
      </InlineField>

      <InlineField label="Endpoint" labelWidth={18}>
        <Input
          onChange={onEndpointChange}
          value={jsonData.endpoint || ''}
          placeholder='MaxCompute Endpoint'
          width={40}
        />
      </InlineField>

      <InlineField label="Project" labelWidth={18}>
        <Input
          onChange={onProjectChange}
          value={jsonData.project || ''}
          placeholder='MaxCompute Project'
          width={40}
        />
      </InlineField>

    </div>
  );
}
