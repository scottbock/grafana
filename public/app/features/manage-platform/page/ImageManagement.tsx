import React from 'react';
import { Button, Card, HorizontalGroup, IconButton, useStyles2 } from '@grafana/ui';
import { GrafanaTheme2 } from '@grafana/data';
import { css } from '@emotion/css';

// TODO: Perfect styling

export const ImageManagement = () => {
  const styles = useStyles2(getStyles);

  return (
    <div>
      <div className={styles.buttonsContainer}>
        <Card heading="Raspberry Pi OS Lite">
          <Card.Meta>
            <a key="link2" href="https://www.raspberrypi.com/software/operating-systems/">
              https://www.raspberrypi.com/software/operating-systems/
            </a>
          </Card.Meta>
          <Card.Figure>
            <img
              src={'https://www.raspberrypi.org/app/uploads/2018/03/RPi-Logo-Reg-SCREEN.png'}
              alt="Raspberry Pi Logo"
            />
          </Card.Figure>
          <Card.Actions>
            <Button icon="play">Start</Button>
            <Button icon="square-shape" variant="secondary">
              Stop
            </Button>
          </Card.Actions>
          <Card.SecondaryActions>
            <IconButton key="delete" name="trash-alt" tooltip="Delete this image" />
          </Card.SecondaryActions>
        </Card>
        <div className={styles.buttonsContainer}>
          <HorizontalGroup spacing="lg">
            <Button icon="plus">Add</Button>
          </HorizontalGroup>
        </div>
      </div>
    </div>
  );
};

const getStyles = (theme: GrafanaTheme2) => ({
  buttonsContainer: css`
    margin-top: 20px;
  `,
});
