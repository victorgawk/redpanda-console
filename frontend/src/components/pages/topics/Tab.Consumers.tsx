import React, { Component } from "react";
import { TopicDetail } from "../../../state/restInterfaces";
import {
    Table,
    Skeleton
} from "antd";
import { observer } from "mobx-react";

import "../../../utils/arrayExtensions";

import { api } from "../../../state/backendApi";
import { motion } from "framer-motion";
import { animProps } from "../../../utils/animationProps";
import { appGlobal } from "../../../state/appGlobal";
import { sortField, makePaginationConfig } from "../../misc/common";
import { uiState } from "../../../state/uiState";
import { DefaultSkeleton } from "../../../utils/tsxUtils";

@observer
export class TopicConsumers extends Component<{ topic: TopicDetail }> {

    pageConfig = makePaginationConfig(20);

    render() {
        const consumers = api.topicConsumers.get(this.props.topic.topicName);
        if (!consumers) return DefaultSkeleton;

        return <div>
            <Table
                style={{ margin: '0', padding: '0' }} size='middle'
                showSorterTooltip={false}
                onRow={(record) =>
                ({
                    onClick: () => appGlobal.history.push('/groups/' + record.groupId),
                })}
                pagination={this.pageConfig}
                onChange={(pagination) => {
                    if (pagination.pageSize) uiState.topicSettings.consumerPageSize = pagination.pageSize;
                    this.pageConfig.current = pagination.current;
                    this.pageConfig.pageSize = pagination.pageSize;
                }}
                rowClassName={() => 'hoverLink'}
                dataSource={consumers}
                rowKey={x => x.groupId}
                columns={[
                    { width: 1, title: 'Group', dataIndex: 'groupId', sorter: sortField('groupId'), defaultSortOrder: 'ascend' },
                    { title: 'Lag', dataIndex: 'summedLag', sorter: sortField('summedLag') },
                ]} />
        </div>
    }
}


