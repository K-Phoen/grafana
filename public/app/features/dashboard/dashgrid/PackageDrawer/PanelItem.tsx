// @ts-nocheck
import { css } from '@emotion/css';
import React, { useState } from 'react';

import type { GrafanaTheme2 } from '@grafana/data';
import { IconButton, Text, useStyles2 } from '@grafana/ui';

export const prettifyTitle = (name: string) =>
  name
    .split('-')
    .filter(Boolean)
    .map((g) => g[0].toUpperCase() + g.slice(1))
    .join(' ');

export const PanelItem = ({ panel, onAddPanel, onRemovePanel }) => {
  const styles = useStyles2(getStyles);
  const [isAdded, setIsAdded] = useState(false);

  return (
    <div className={styles.panelSection}>
      {isAdded ? (
        <IconButton
          name="times-circle"
          size="xl"
          variant="destructive"
          className={styles.button}
          onClick={() => {
            onRemovePanel(panel);
            setIsAdded(false);
          }}
        />
      ) : (
        <IconButton
          name="plus-circle"
          size="xl"
          variant="primary"
          className={styles.button}
          onClick={() => {
            onAddPanel(panel);
            setIsAdded(true);
          }}
        />
      )}
      <div className={styles.panelTextSection}>
        <Text element="h6" color="maxContrast" weight="light">
          {prettifyTitle(panel.metadata.name)}
        </Text>
        <Text element="p" color="secondary" italic>
          {panel.metadata.description || 'No description'}
        </Text>
      </div>
    </div>
  );
};

function getStyles(theme: GrafanaTheme2) {
  return {
    button: css({
      margin: '0px 12px',
    }),
    panelSection: css({
      display: 'flex',
      flexDirection: 'row',
      marginBottom: '8px',
    }),
    panelTextSection: css({
      display: 'flex',
      flexDirection: 'column',
    }),
  };
}
