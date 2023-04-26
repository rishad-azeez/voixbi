import { DataFrame, DataFrameJSON, dataFrameToJSON } from '@grafana/data';
import appEvents from 'app/core/app_events';
import { GRAFANA_DATASOURCE_NAME } from 'app/features/alerting/unified/utils/datasource';
import { PanelModel } from 'app/features/dashboard/state';
import { ShowConfirmModalEvent } from 'app/types/events';

import { GrafanaQuery, GrafanaQueryType } from './types';

/**
 * Will show a confirm modal if the current panel does not have a snapshot query.
 * If the confirm modal is shown, and the user aborts the promise will resolve with a false value,
 * otherwise it will resolve with a true value.
 */
export function onUpdatePanelSnapshotData(panel: PanelModel, frames: DataFrame[]): Promise<boolean> {
  return new Promise<boolean>((resolve) => {
    if (panel.datasource?.uid === GRAFANA_DATASOURCE_NAME) {
      updateSnapshotData(frames, panel);
      resolve(true);
      return;
    }

    appEvents.publish(
      new ShowConfirmModalEvent({
        title: 'Change to panel embedded data',
        text: 'If you want to change the data shown in this panel Grafana will need to remove the panels current query and replace it with a snapshot of the current data. This enabled you to edit the data',
        yesText: 'Continue',
        icon: 'pen',
        onConfirm: () => {
          updateSnapshotData(frames, panel);
          resolve(true);
        },
        onDismiss: () => {
          resolve(false);
        },
      })
    );
  });
}

function updateSnapshotData(frames: DataFrame[], panel: PanelModel) {
  const snapshot: DataFrameJSON[] = frames.map((f) => dataFrameToJSON(f));

  const query: GrafanaQuery = {
    refId: 'A',
    queryType: GrafanaQueryType.Snapshot,
    snapshot,
    datasource: { uid: GRAFANA_DATASOURCE_NAME },
  };

  panel.updateQueries({
    dataSource: { uid: GRAFANA_DATASOURCE_NAME },
    queries: [query],
  });

  panel.refresh();
}
